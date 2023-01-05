# all packages in ~/go/pkg/mod/...

# create package

```bash
$ mkdir ~/go/src/aPackage
$ cp aPackage.go ~/go/src/aPackage/
$ go install aPackage
$ cd ~/go/pkg/darwin_amd64/
$ ls -l aPackage.a
-rw-r--r-- 1 mtsouk staff 4980 Dec 22 06:12 aPackage.a
```

# module version

## create module

```bash
$ git clone git@github.com:mactsouk/myModule.git
$ go mod init
go: creating new go.mod: module github.com/mactsouk/myModule
$ touch myModule.go
$ vi myModule.go
$ git add .
$ git commit -a -m "Initial version 1.0.0"
$ git push
$ git tag v1.0.0
$ git push -q origin v1.0.0
$ go list
github.com/mactsouk/myModule
$ go list -m
github.com/mactsouk/myModule
$ cat go.mod
module github.com/mactsouk/myModule
go 1.12

```

## using module v1

### run

```bash
$ export GO111MODULE=on
$ go run useModule.go
go: finding github.com/mactsouk/myModule v1.0.0
go: downloading github.com/mactsouk/myModule v1.0.0
go: extracting github.com/mactsouk/myModule v1.0.0
Version 1.0.0

```

### build

```bash
$ go mod init hello
go: creating new go.mod: module hello
$ go build
$ cat go.sum
github.com/mactsouk/myModule v1.0.0
h1:eTCn2Jewnajw0REKONrVhHmeDEJ0Q5TAZ0xsSbh8kFs=
github.com/mactsouk/myModule v1.0.0/go.mod
h1:s3ziarTDDvaXaHWYYOf/ULi97aoBd6JfnvAkM8rSuzg=
$ cat go.mod
module hello
go 1.12
require github.com/mactsouk/myModule v1.0.0
```

## create module v1.1

```bash
$ vi myModule.go
$ git commit -a -m "v1.1.0"
[master ddd0742] v1.1.0
    1 file changed, 1 insertion(+), 1 deletion(-)
$ git push
$ git tag v1.1.0
$ git push -q origin v1.1.0
```

## using module v1.1 in doker

```bash
$ docker run --rm -it golang:latest
root@884c0d188694:/go# cd /tmp
root@58c5688e3ee0:/tmp# go version
go version go1.13 linux/amd64
$ root@58c5688e3ee0:/tmp# ls -l
total 4
-rw-r--r-- 1 root root 91 Mar 2 19:59 useUpdatedModule.go
$root@58c5688e3ee0:/tmp# export GO111MODULE=on
$root@58c5688e3ee0:/tmp# go run useUpdatedModule.go
go: finding github.com/mactsouk/myModule v1.1.0
go: downloading github.com/mactsouk/myModule v1.1.0
go: extracting github.com/mactsouk/myModule v1.1.0
Version 1.1.0
```

## create module v2

```bash
$ git checkout -b v2
Switched to a new branch 'v2'
$ git push --set-upstream origin v2
$ vi go.mod
$ cat go.mod
module github.com/mactsouk/myModule/v2
go 1.12
$ git commit -a -m "Using 2.0.0"
[v2 5af2269] Using 2.0.0
2 files changed, 2 insertions(+), 2 deletions(-)
$ git tag v2.0.0
$ git push --tags origin v2
Counting objects: 4, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 441 bytes | 441.00 KiB/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
To github.com:mactsouk/myModule.git
  * [new branch]     v2 -> v2
  * [new tag]        v2.0.0 -> v2.0.0
$ git --no-pager branch -a
    master
* v2
  remotes/origin/HEAD -> origin/master
  remotes/origin/master
  remotes/origin/v2

```

## using module v2

```bash
$ docker run --rm -it golang:latest
$ root@191d84fc5571:/go# cd /tmp
$ root@191d84fc5571:/tmp# cat > useV2.go
package main
import (
    v "github.com/mactsouk/myModule/v2"
)
func main() {
    v.Version()
}
$ root@191d84fc5571:/tmp# export GO111MODULE=on
$ root@191d84fc5571:/tmp# go run useV2.go
go: finding github.com/mactsouk/myModule/v2 v2.0.0
go: downloading github.com/mactsouk/myModule/v2 v2.0.0
go: extracting github.com/mactsouk/myModule/v2 v2.0.0
Version 2.0.0
```

## create module v2.1

```bash
$ vi myModule.go
$ git commit -a -m "v2.1.0"
$ git push
$ git tag v2.1.0
$ git push -q origin v2.1.0
```

## using module v2.1

```bash
$ docker run --rm -it golang:1.12
$ root@ccfcd675e333:/go# cd /tmp/
$ root@ccfcd675e333:/tmp# cat > useUpdatedV2.go
package main
import (
    v "github.com/mactsouk/myModule/v2"
)
func main() {
    v.Version()
}
$ root@ccfcd675e333:/tmp# ls -l
total 4
-rw-r--r-- 1 root root 92 Mar 2 20:34 useUpdatedV2.go
$ root@ccfcd675e333:/tmp# go run useUpdatedV2.go
useUpdatedV2.go:4:2: cannot find package "github.com/mactsouk/myModule/v2" in any of:
    /usr/local/go/src/github.com/mactsouk/myModule/v2 (from $GOROOT)
    /go/src/github.com/mactsouk/myModule/v2 (from $GOPATH)
$ root@ccfcd675e333:/tmp# export GO111MODULE=on
$ root@ccfcd675e333:/tmp# go run useUpdatedV2.go
go: finding github.com/mactsouk/myModule/v2 v2.1.0
go: downloading github.com/mactsouk/myModule/v2 v2.1.0
go: extracting github.com/mactsouk/myModule/v2 v2.1.0
Version 2.1.0
```

# go mod vendor

```bash
$ cd useTwoVersions
$ go mod init useV1V2
go: creating new go.mod: module useV1V2
$ go mod vendor

```
