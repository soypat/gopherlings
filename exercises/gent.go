// gent is a tool to facilitate numeration of
// gopherling exercises.
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		usage("need command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "insert":
		if len(os.Args) < 3 {
			usage("need argument for insert")
			os.Exit(1)
		}
		commandInsert(os.Args[2])
	case "collapse":
		commandCollapse()
	default:
		usage("unknown command " + os.Args[1])
		os.Exit(1)
	}
	log.Println("finished succesfully")
}

func commandCollapse() {
	exercises, err := parseDirExercises(".")
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(byNumber(exercises))
	for i, ex := range exercises {
		expectedNum := i + 1
		if ex.num == expectedNum {
			continue
		}
		renamed := e{num: i + 1, name: ex.name}
		if _, err := os.Stat(renamed.dirName()); err == nil {
			log.Fatalf("can't rename %q to already existing %q", ex.dirName(), renamed.dirName())
		}
		err := os.Rename(ex.dirName(), renamed.dirName())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("collapse: renaming %q -> %q\n", ex.dirName(), renamed.dirName())
	}
}

func commandInsert(name string) {
	newExercise, err := parseExerciseDirName(name)
	if err != nil {
		log.Fatal("parsing argument exercise name: ", err)
	}
	if _, err := os.Stat(name); err == nil {
		log.Fatalf("exercise %s already exists", name)
	}
	log.Printf("inserting exercise %q", newExercise.dirName())
	oldExercises, err := parseDirExercises(".")
	if err != nil {
		log.Fatal(err)
	}
	for i := range oldExercises {
		if oldExercises[i].name == newExercise.name {
			log.Fatalf("exercise %q already exists", newExercise.name)
		}
	}
	// Modify old exercise names and add new exercise.
	for _, exercise := range oldExercises {
		if exercise.num < newExercise.num {
			continue
		}
		renamedExercise := e{
			num:  exercise.num + 1.,
			name: exercise.name,
		}
		err := os.Rename(exercise.dirName(), renamedExercise.dirName())
		if err != nil {
			log.Fatal("")
		}
	}
	err = os.Mkdir(name, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

type e struct {
	num  int
	name string
}

func parseExerciseDirName(name string) (e, error) {
	if len(name) < 5 {
		return e{}, errors.New("exercise name too short")
	}
	n, err := strconv.Atoi(name[0:3])
	if err != nil {
		return e{}, err
	}
	en := e{num: n, name: name[4:]}
	if en.dirName() != name {
		return e{}, fmt.Errorf("generated directory name %q does not match argument name %q", en.dirName(), name)
	}
	return en, nil
}

func (en e) dirName() string {
	return fmt.Sprintf("%03d-%s", en.num, en.name)
}

var excludeDir = []string{}

func isExcluded(s string) bool {
	for _, ex := range excludeDir {
		if ex == s {
			return true
		}
	}
	return false
}

func parseDirExercises(dir string) ([]e, error) {
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}
	var found []e
	for _, entry := range entries {
		dirName := entry.Name()
		if !entry.IsDir() || isExcluded(dirName) {
			continue
		}
		exercise, err := parseExerciseDirName(dirName)
		if err != nil {
			return nil, errors.New("parsing existing exercise directory: " + err.Error())
		}
		newName := (e).dirName(e{num: exercise.num + 1, name: exercise.name})
		_, err = os.Stat(newName)
		if err == nil {
			fmt.Errorf("conflict rename: exercise rename %q already exists", newName)
		}
		found = append(found, exercise)
	}
	return found, nil
}

func usage(msg string) {
	fmt.Fprintf(os.Stderr, "gopherling exercise numeration tool (gent).\n")
	fmt.Fprintf(os.Stderr, "usage:\n\tgent <command> [argument]\n")
	fmt.Fprintf(os.Stderr, "Supported commands:\n")
	fmt.Fprintf(os.Stderr, "\tinsert [name of exercise directory]\n")
	fmt.Fprintf(os.Stderr, "inserts an exercise directory and renames existing exercise numbers to avoid number collisions\n")
	fmt.Fprintf(os.Stderr, "\trenames exercise numbers so that exercise numbers increase by 1\n")
	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
}

type byNumber []e

func (a byNumber) Len() int           { return len(a) }
func (a byNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byNumber) Less(i, j int) bool { return a[i].num < a[j].num }
