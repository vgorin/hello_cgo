hello_cgo
=========

The simplest cgo demo, describes and hopefully solves my old problem "Weird things happening while compiling a Go program with C code (go build, Eclipse, Mac OS X)"


-------------------------
https://groups.google.com/forum/#!searchin/Golang-Nuts/vgorin/golang-nuts/IxPSrh1X7QE/JnWfb8vMX5MJ


Original problem post(s):
=========================

Hello, guys,

I'm compiling a simple Go program which just executes some C code from the *.c files.
All program files are located in the same directory, these are five *.c files, five *.h files (headers are used in *.c files) and one *.go file with package main and function main declared in it.

I'm compiling the app on OS X 10.8.4.
When I compile with Eclipse I do not have any problems - the program builds and runs perfectly.

Now I'm trying to build with "go build" and receive a weird error. The error says that there are "69 duplicate symbols for architecture x86_64". The files which have errors are _obj/*.cgo2.o and _obj/*.o - I think it means that this is a linker error.
The thing I do not understand is how then Eclipse compiles the code with no errors? Does it use mechanism other then go build? It definitely uses the same gcc, as there is no gcc path specified in Eclipse itself so it can use just the one which in the $PATH.

Please post your ideas why such a thing can happen.

Thank you

=========================
To simplify everything I've created a more simple app with only two files: hello.go and hello.c:

hello.c:

#include <stdlib.h>

int rnd() {
	return random();
}


hello.go:

package main

// #include "hello.c"
import "C"

import "fmt"

func Random() int {
	return int(C.rnd())
}

func main() {
	fmt.Println(Random());
}

The symptoms of compiling these are exactly the same. It compiles in eclipse, but not with go build:

KIEV-AIR:src vgorin$ go build
# _/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src
duplicate symbol _rnd in:
    $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.cgo2.o
    $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.o
ld: 1 duplicate symbol for architecture x86_64
collect2: ld returned 1 exit status

Output of go build -x:

KIEV-AIR:src vgorin$ go build -x
WORK=/var/folders/j1/5yp9qqm5429_j05vcg_gfd0m0000gn/T/go-build388081714
mkdir -p $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/
cd /Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src
/usr/local/go/pkg/tool/darwin_amd64/cgo -objdir $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/ -- -I $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/ hello.go
/usr/local/go/pkg/tool/darwin_amd64/6c -F -V -w -I $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/ -I /usr/local/go/pkg/darwin_amd64 -o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_defun.6 -D GOOS_darwin -D GOARCH_amd64 $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_defun.c
gcc -I . -g -O2 -fPIC -m64 -pthread -fno-common -print-libgcc-file-name
gcc -I . -g -O2 -fPIC -m64 -pthread -fno-common -I $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/ -o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_main.o -c $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_main.c
gcc -I . -g -O2 -fPIC -m64 -pthread -fno-common -I $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/ -o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_export.o -c $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_export.c
gcc -I . -g -O2 -fPIC -m64 -pthread -fno-common -I $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/ -o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.cgo2.o -c $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.cgo2.c
gcc -I . -g -O2 -fPIC -m64 -pthread -fno-common -I $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/ -o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.o -c ./hello.c
gcc -I . -g -O2 -fPIC -m64 -pthread -fno-common -o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_.o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_main.o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/_cgo_export.o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.cgo2.o $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.o
# _/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src
duplicate symbol _rnd in:
    $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.cgo2.o
    $WORK/_/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src/_obj/hello.o
ld: 1 duplicate symbol for architecture x86_64
collect2: ld returned 1 exit status

Environment variables are set:
KIEV-AIR:src vgorin$ echo $GOROOT
/usr/local/go
KIEV-AIR:src vgorin$ echo $GOPATH
/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/
KIEV-AIR:src vgorin$ echo $GOBIN
/usr/local/go/bin
KIEV-AIR:src vgorin$ echo $GOOS
darwin
KIEV-AIR:src vgorin$ echo $GOARCH
amd64
KIEV-AIR:src vgorin$ pwd
/Users/vgorin/DEVELOP/eclipse_workspace/hello_world_c/src
=========================
My considerations as for 23/10/2013

'go build' tries to compile every file it finds (both *.go and *.c types) and then it links everything together.
'go build gofile.go' tries to compile only gofile.go and its dependencies.
Thus, for the first command to work including *.c files into go code with // #include won't work, need to include header files *.h only, corresponding *.c files will be compiled separetelly and then linked;
for the second command to work we need to include *.c files or provide already compiled libraries if we want to include header files *.h

This project shows the approach which works, but perhaps its not ideal – probably there is a better solution with ompiler and linker flags – this should be investigated.
