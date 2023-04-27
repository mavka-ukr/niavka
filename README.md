```shell
go run main.go тест.н >> hello.ll
llc -filetype=obj hello.ll -o hello.o
clang hello.o -o hello
./hello
```
