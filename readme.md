## git-pair

<a href='http://www.recurse.com' title='Made with love at the Recurse Center'><img src='https://cloud.githubusercontent.com/assets/2883345/11325206/336ea5f4-9150-11e5-9e90-d86ad31993d8.png' height='20px'/></a>

> Make it easy to add co-authors to your project

```
$ go install tab.go
```

Make sure that your `go/bin` directory is in your path. You can check this by `echo $PATH`. If the go directory is not there then go into your zsh or bash file and add that.

Add the following to your `.zshrc` file:

```
PROG=tab source $PATH/to/autocomplete/bash_autocomplete
```
