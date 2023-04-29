```shell
go run main.go привіт.н >привіт.ll
llc -filetype=obj привіт.ll -o привіт.o
clang привіт.o -o привіт

./привіт
```
або
```shell
./compile.sh привіт.н привіт

./привіт
```
