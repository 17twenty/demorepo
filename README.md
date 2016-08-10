# demorepo

## Purpose ##

Barebones setup to show how a binary can be ```go getted``` and build and installed automatically.

## Usage ##
```
go get github.com/17twenty/demorepo/...
```

As long as your $GOPATH/bin/ is in your $PATH you should be able to execute the binary as:

```
$ newbin 
2016/08/10 15:53:36 Here I am
```

That's it.

## Further Information ##

You can make as many bins available as you want with a layout like this:
```
github.com/17twenty/
  cmd/
    s3get/
      main.go
    s3mount/
      main.go
    s3put/
      main.go
    newbin/
      main.go
```

If you want to build a library, nest your main app and put consumers of it in subdirs like this:
```
github.com/17twenty/
  epics3lib.go
  amazingdep.go
  cmd/
    s3get/
      main.go
    s3mount/
      main.go
    s3put/
      main.go
```

Then your user can use ```go get``` to install the library, and ```go get github.com/17twenty/demorepo/...``` to fetch, build and install the binaries as well.

See [this article](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091#.zag9ikesn) for more.
