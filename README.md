# gopherlings
📘️ Learn Go by fixing tiny incorrect programs

This project was directly inspired by the great [ziglings](https://github.com/ratfactor/ziglings) project which itself was inspired by [rustlings](https://github.com/rust-lang/rustlings).

For a first time learner it is suggested you complement this material with another source such as

- [Go by example](https://gobyexample.com/). A lot of material here is based on this project!
- [Go go-to guide](https://yourbasic.org/golang/) by yourbasic. Great source for beginners.

## Intended Audience
These exercises will probably be difficult if you've never programmed before. 

The exercises should be self-contained and self-explained, though this is a WIP
and suggestions are welcome!

## Instructions: Running exercises
Requires a [Go installation](https://go.dev/dl/) to run the examples.
### Instructions using VSCode

1. Download the repository (or alternatively clone it)

2. Install the VSCode **Go** extension authored by _Go Team at Google_

3. Open the `gopherlings` folder in VSCode

4. Navigate to the exercise file, i.e. [`exercises/001-hello/hello.go`](exercises/001-hello/hello.go)

5. Once the `hello.go` file is open you may edit it and press <kbd>F5</kbd> to run it. Output will be shown in the Debug Console.

### Instructions using terminal
1. Clone repository
    ```sh
    git clone https://github.com/soypat/gopherlings.git 
    ```

2. Navigate to example's directory
    ```sh
    cd gopherlings/exercises/001-hello
    ```

3. Edit the file so it is correct and run it with `go run`
    ```sh
    go run hello.go
    ```
   Optionally, instead of step 3, use `gopherlings watch` to auto reload your code after you save your code. Then you will see your code output without any action. You can simply focus on your code.
   ```sh
   go install github.com/soypat/gopherlings/cmd/gopherlings@latest
   gopherlings watch
   gopherlings help  # To see other commands
   ```
   Note: You need to define your `GOBIN` environment variable before running `go install` and add it to your `PATH`. `GOBIN` specifies the directory to which go binaries are installed to.

## Roadmap
#### Short term
* Add quizzes after $n$ exercises.
  - Possibly add several quizzes of varying difficulty inside quiz directory.
* Improve the helper program.
  - A `hint` subcommand or interactive hints during the `watch` command.
  - Allow `run` to take an argument that runs a specific exercise.
  - Add a subcommand to run all exercises and check them against expected output.
  - Catch errors that prevent the helper from running exercises (like an incomplete Go installation).

#### Long term
* Have exercises that cover the entire [Go spec](https://go.dev/ref/spec).
