<h1 style="text-align:center">
    Go Study Note
</h1>
<p style="text-align:right">
    <b>Axect</b>
</p>

## 1. Start Setting
1. Install Golang
2. Setting Path (GOROOT, GOPATH)
    > You can check setting by go env  

3. Install VSCode and install go extension
4. Install All
5. Go to GOPATH, mkdir src/github.com/username/your_repository_name
6. Enjoy Go (If you use windows - terminal setting : Powershell or use linux then use ordinary terminal)

## 2. Basic Structure
1. In your repository folder, make custom folder
2. Make custom go file as same as your folder name
3. And input as follows.
    > ```Go
    > package file_name
    > 
    > // Capital_Letters is blabla
    > func Capital_Letters() {
    >   Write_your_Code
    > }
    >```

4. Make cmd folder
5. Make main go file :
    > ```Go
    > package main
    > import (
    >   "github.com/username/repository_name/filename"
    >)
    > 
    > func main() {
    >   filename.Capital_Letters()
    >}
    >```

6. There are two things to compile and run go.
    1. Simple : go run main.go
    2. Stable : go build main.go
        * Windows : .\main.exe
        * Linux : ./main
* Don't forget 'add, commit, push, pull'
