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

5. One the `hello.go` file is open you may edit it and press <kbd>F5</kbd> to run it. Output will be shown in the Debug Console.

### Instructions using terminal
1. Clone repository
    ```sh
    git clone https://github.com/soypat/gopherlings.git 
    ```

2. Navigate to example's directory
    ```sh
    cd gopherlings/exercise/001-hello
    ```

3. Edit the file so it is correct and run it with `go run`
    ```sh
    go run hello.go
    ```