
Slides are here : https://docs.google.com/presentation/d/1A_4_oWK5BohgJeYAdbDw1rPx-vJerKxBYV2QD7kw0k0/edit?usp=sharing

## Before going further, set your GOPATH system variable as this folder
$ export GOPATH=$(pwd)

## To compile a go single file progam :
$ go build [xxx].go

## To compie a go package and make it accessible in your local environment :
$ cd [path of the folder containing the package]
$ go install

## To build and run a go program :
$ go run [xxx].go

## To generate the package documentation
$ godoc [path of the folder containing the package] [function name]

## Get a package from internet, compile it and store it in our local repository
$ go get [github.com/go-chef/chef]

## To test a package
$ go test [github.com/go-chef/chef]
