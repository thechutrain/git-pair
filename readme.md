## tab

> using tab completion in a separate project

```
$ go install tab.go
```

Make sure that your `go/bin` directory is in your path. You can check this by `echo $PATH`. If the go directory is not there then go into your zsh or bash file and add that.

Add the following to your `.zshrc` file:

```
PROG=tab source $PATH/to/autocomplete/bash_autocomplete
```
