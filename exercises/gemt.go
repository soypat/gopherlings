package main

import (
	"errors"
	"fmt"
	"log"
	"os"
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
	default:
		usage("unknown command " + os.Args[1])
		os.Exit(1)
	}
	log.Println("finished succesfully")
}

func commandInsert(name string) {
	newExercise, err := parseExerciseName(name)
	if err != nil {
		log.Fatal("parsing argument exercise name: ", err)
	}
	log.Printf("inserting exercise %q", newExercise.dirName())
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	var oldExercises []e
	for _, entry := range entries {
		dirName := entry.Name()
		if !entry.IsDir() || isExcluded(dirName) {
			continue
		}
		exercise, err := parseExerciseName(dirName)
		if err != nil {
			log.Fatal("parsing existing exercise directory: ", err)
		}
		newName := (e).dirName(e{num: exercise.num + 1, name: exercise.name})
		_, err = os.Stat(newName)
		if err == nil || newName == newExercise.dirName() {
			log.Fatalf("conflict rename: exercise rename %q already exists or coincides with insertion", newName)
		}
		oldExercises = append(oldExercises, exercise)
	}

	// Modify old exercise names and add new exercise.
	for _, exercise := range oldExercises {
		newExercise := e{
			num:  exercise.num + 1.,
			name: exercise.name,
		}
		err := os.Rename(exercise.dirName(), newExercise.dirName())
		if err != nil {
			log.Fatal("")
		}
	}
	err = os.Mkdir(name, os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
}

type e struct {
	num  int
	name string
}

func parseExerciseName(name string) (e, error) {
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

func usage(msg string) {
	fmt.Fprintf(os.Stderr, "gopherling exercise manipulation tool (gemt)\n")
	fmt.Fprintf(os.Stderr, "usage:\n\tgemt <command> [argument]\n")
	fmt.Fprintf(os.Stderr, "Supported commands:\n")
	fmt.Fprintf(os.Stderr, "\tinsert [name of exercise directory]\n")
	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
}
