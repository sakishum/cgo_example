g++ foo.cpp -fPIC -shared -o libfoo.so
go build foo.go
LD_LIBRARY_PATH=. ./foo
rm -f foo libfoo.so
