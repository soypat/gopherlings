package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

const exercisesDir = "exercises"

// maxRuntime is the maximum time an exercise is allowed to run in watch mode.
const maxRuntime = 5 * time.Second

var reStillGoing = regexp.MustCompile(`(?m)^\s*//\s+I\s+AM\s+STILL\s+GOING`)
var reExercise = regexp.MustCompile(`^(?P<id>\d+)-(?P<name>.+)$`)

var plain = color.New()
var red = color.New(color.FgRed)
var yellow = color.New(color.FgYellow)
var green = color.New(color.FgGreen)

var rootCmd = &cobra.Command{
	Use:   "gopherlings",
	Short: `Learn Go by fixing tiny incorrect programs`,
	Long:  rootHelp,
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(watchCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		red.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// errExit quits with friendly error messages from within Cobra commands. If
// err is nil, this returns without printing anything or exiting.
func errExit(err error) {
	if err == nil {
		// No error, continue!
		return
	}

	if errors.Is(err, ErrAllDone) {
		// Error: success ;)
		plain.Printf(goBuildSomething)
		os.Exit(0)
	} else if errors.Is(err, ErrNoExercisesDir) {
		red.Fprintln(os.Stderr, `No "exercises" directory.`)
		red.Fprintln(os.Stderr, "Please run this from the root of a Gopherlings repo checkout.")
	} else {
		red.Fprintln(os.Stderr, err)
	}

	os.Exit(1)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all exercises with their status",
	Run: func(cmd *cobra.Command, args []string) {
		exs, err := listExercises(exercisesDir)
		errExit(err)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		for _, ex := range exs {
			var status string
			stillGoing, err := ex.StillGoing()
			if err != nil {
				status = fmt.Sprintf("error: %s", err)
			} else if stillGoing {
				status = "still-going"
			} else {
				status = "done"
			}
			name := filepath.Base(ex.Dir)
			fmt.Fprintln(w, name, "\t", status)
		}
		w.Flush()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the next unfinished exercise once",
	Run: func(cmd *cobra.Command, args []string) {
		exs, err := listExercises(exercisesDir)
		errExit(err)

		next, err := nextExercise(exs)
		errExit(err)

		next.Run(context.Background())
	},
}

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: `Continuously run the next unfinished exercise`,
	Run: func(cmd *cobra.Command, args []string) {
		exs, err := listExercises(exercisesDir)
		errExit(err)

		err = watch(exs)
		errExit(err)
	},
}

var ErrNoExercisesDir = errors.New(`no "exercises" directory`)

func listExercises(dir string) ([]Exercise, error) {
	subdirs, err := os.ReadDir(dir)

	if errors.Is(err, fs.ErrNotExist) {
		return nil, ErrNoExercisesDir
	}
	if err != nil {
		return nil, err
	}

	var exs []Exercise

	for _, sd := range subdirs {
		name := sd.Name()

		m := reExercise.FindStringSubmatch(name)
		if len(m) == 0 {
			// This is likely a helper of some kind. Skip it.
			continue
		}

		reldir := filepath.Join(dir, name)
		exs = append(exs, Exercise{reldir})
	}

	return exs, nil
}

var ErrAllDone = errors.New("all exercises compelete")

func nextExercise(exs []Exercise) (*Exercise, error) {
	for _, ex := range exs {
		ex := ex
		stillGoing, err := ex.StillGoing()
		if err != nil {
			return nil, err
		}

		if stillGoing {
			return &ex, nil
		}
	}

	// All done!
	return nil, ErrAllDone
}

// watch continuously runs the next unfinished exercise. This blocks until all
// exercises are complete or there's an error finding the next exercise.
func watch(exs []Exercise) error {
	for {
		next, err := nextExercise(exs)
		if err != nil {
			return err
		}

		clear()

		ctx, cancel := context.WithTimeout(context.Background(), maxRuntime)

		done := make(chan struct{})
		go func() {
			next.Run(ctx)
			close(done)
		}()

		// Immediately cancel the run context as soon as a change or error
		// occurs. To avoid interleaving output, wait for the exercise process
		// to exit before continuing.
		err = waitForChanges()
		cancel()
		<-done
		if err != nil {
			return err
		}
	}
}

// clear prints the escape sequence to reset the terminal.
func clear() {
	fmt.Printf("\x1bc")
}

// waitForChanges watches an "exercises" directory for changes and blocks until
// the first event or watcher error.
func waitForChanges() error {
	dir := exercisesDir

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	changed := make(chan error)
	go func() {
		select {
		case _ = <-watcher.Events:
			changed <- nil
		case err := <-watcher.Errors:
			changed <- err
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		return err
	}

	subdirs, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// fsnotify doesn't automatically recursively watch directories. There
	// should be exactly one level of nested directories underneath the
	// exercise root, so this is enough to watch everything.
	for _, subdir := range subdirs {
		path := filepath.Join(dir, subdir.Name())
		err = watcher.Add(path)
		if err != nil {
			return err
		}
	}

	return <-changed
}

// Exercise represents an exercise's relative location on the filesystem.
type Exercise struct {
	Dir string
}

// StillGoing looks for the "I AM STILL GOING" comment within an exercise
// directory to determine whether the exercise is still ongoing.
//
// This seems really optimistic, but it can't rely on running the exercise.
// Some unfinished exercises compile and run without errors. And some of those
// don't terminate at all!
func (e Exercise) StillGoing() (bool, error) {
	files, err := os.ReadDir(e.Dir)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		name := file.Name()

		// Skip non-Go files like editor swap files.
		if filepath.Ext(name) != ".go" {
			continue
		}

		path := filepath.Join(e.Dir, name)

		f, err := os.Open(path)
		if err != nil {
			return false, err
		}
		defer f.Close()

		r := bufio.NewReader(f)
		if reStillGoing.MatchReader(r) {
			return true, nil
		}
	}

	// No file in the package has the comment, so it must be done!
	return false, nil
}

// Run compiles and executes an exercise.
func (e Exercise) Run(ctx context.Context) {
	// `go build` requires the dot+separator prefix for local packages.
	pkgPath := "." + string(filepath.Separator) + e.Dir

	err := runPkg(ctx, pkgPath)

	switch {
	case err == nil:
		green.Println()
		green.Println("Successfully ran exercise!")
		green.Println(`Remove the "I AM STILL GOING" comment to continue to the next exercise.`)
		green.Println()

	case errors.Is(err, context.Canceled):
		// Explicit cancel from the exercise driver (run or watch). Let the
		// canceler decide what to print, if anything.

	case errors.Is(err, context.DeadlineExceeded):
		red.Println("The exercise ran for too long without exiting.")
		red.Println("Exercises should generally finish running very quickly.")
		red.Println("")
		red.Println("Is there something preventing the program from exiting?")

	case isSignal(err):
		// Killed by signal for some other reason. This is probably not related
		// to the exercise.
		yellow.Println("The exercise stopped unexpectedly.")
		yellow.Println("")
		yellow.Println(err)

	default:
		// A real compile or runtime error related to the exercise.
		red.Println(err)
	}
}

func runPkg(ctx context.Context, pkgPath string) error {
	// Compile and run separately for two reasons:
	//
	// 1. It's easier to separate compile errors and run output.
	// 2. `go run ./pkg` runs the compiled command in a subprocess, and there
	//    isn't an obvious cross-platform way to make exec.Cmd kill that
	//    subprocess if the context is canceled.

	plain.Println("Compiling", pkgPath)
	f, err := compilePkg(ctx, pkgPath)
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())

	plain.Println("Executing", pkgPath)
	plain.Println()
	return execFile(ctx, f)
}

func compilePkg(ctx context.Context, pkg string) (*os.File, error) {
	name := filepath.Base(pkg)
	f, err := os.CreateTemp("", name)
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(ctx, "go", "build", "-o", f.Name(), pkg)

	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	err = cmd.Run()
	if err != nil {
		os.Remove(f.Name())

		plain.Println("Compilation failed.")
		plain.Println()

		red.Println(stderr)
		return nil, err
	}

	return f, nil
}

func execFile(ctx context.Context, f *os.File) error {
	cmd := exec.CommandContext(ctx, f.Name())

	// TODO: Use io.LimitReader to prevent excessive output.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	// If both the context error and the process error are non-nil, cmd.Run()
	// prefers the process error.
	//
	// But we have the opposite preference! If the exercise driver canceled the
	// context, the way the process died isn't actually meaningful.
	ctxErr := ctx.Err()
	if ctxErr != nil {
		return ctxErr
	}

	return err
}

func isSignal(err error) bool {
	var exit *exec.ExitError
	if errors.As(err, &exit) {
		if exit.ExitCode() == -1 {
			return true
		}
	}
	return false
}

const rootHelp string = `Gopherlings is designed to help beginners learn the basics of Go.

The Gopherlings exercises are small incomplete programs that will teach you Go
concepts as you solve them. Each one will contain a syntax error or logic error
for you to find and fix.

The Gopherlings application (this one!) will help you run these exercises to
give you a fast feedback loop for learning. Make sure to run this from within a
checked-out Gopherlings repo with an "exercises" directory.

To mark an exercise as completed, remove the "I AM STILL GOING" comment near
the top of the file. Many exercises have more than one solution, so this lets
you continue trying things on the same exercise even if it runs to completion.

To get started in watch mode, run ` + "`gopherlings watch`" + ` and look for the
first error message.
`

const goBuildSomething = `You've completed all the Gopherlings exercises!

We hope you had fun and learned something along the way.

If you noticed any issues, please report them to the repo.
You can also contribute exercises or ideas for exercises.

https://github.com/soypat/gopherlings

And with that...

    go buildSomething()

`
