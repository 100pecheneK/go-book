# Build C

```bash
$ gcc -o main main.c mainGo.o
$ ./main
```

# Build Go

```bash
$ go build -o mainGo.o -buildmode=c-shared mainGo.go
```

# Run

```bash
$ ./main
```
