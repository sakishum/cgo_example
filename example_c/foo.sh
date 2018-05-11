gcc -c foo.c
ar rv libfoo.a foo.o
go build foo.go
./foo
rm -f foo foo.o libfoo.a
