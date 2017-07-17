
Code snippets written in c/c++/golang


A workspace contains three subdirectories at its root:

src – This directory contains source files organized as packages. You will write your Go applications inside this directory.1
pkg – This directory contains Go package objects.
bin – This directory contains executable programs.
You have to specify the location of the workspace before developing Go programs. The GOPATH environment variable is used for specifying the location of Go workspaces.

The packages from the standard library are available at the “pkg” subdirectory of the GOROOT directory

When you import packages, the Go compiler will look on the locations specified by the environment variable GOROOT and GOPATH

Packages from the standard library are available in the GOROOT location.
The packages that are created by yourself, and third-party packages which you have imported, are available in the GOPATH location.


Init Function
When you write Go packages, you can provide a function “init” that will be called at the beginning of the execution time.

package db
import (
       "gopkg.in/mgo.v2"
               "gopkg.in/mgo.v2/bson"
	       )
func init {
       // initialization



