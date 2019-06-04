## git-pair 🍐🍐

<a href='http://www.recurse.com' title='Made with love at the Recurse Center'><img src='https://cloud.githubusercontent.com/assets/2883345/11325206/336ea5f4-9150-11e5-9e90-d86ad31993d8.png' height='20px'/></a>

> Makes it easy to add co-authors to your project

# Installing

```
$ go install git-pair.go
```

Make sure that your `go/bin` directory is in your path. You can check this by `echo $PATH`. If the go directory is not there then go into your zsh or bash file and add that.

Add the following to your `.zshrc` file:

```
PROG=pair source $PATH/to/autocomplete/bash_autocomplete
```

# Usage

To initiate a pairing session just type:

```
$pair add GITHUB_USERNAME GITHUB_EMAIL
```

To stop pairing with a single user:

```
$ pair remove GITHUB_USERNAME
```

To stop pairing with everyone:

```
$ pair stop
```

To see if you are currently pairing with anyone:

```
$ pair status
```

# Description

This cli uses the `prepare-commit-msg` git hook to append co-authors to your commit messages for you.

Example:

```
wip - git commit message, pairing is fun!

# Added by 🍐
Co-authored-by: pairbot <pairbot@email.com>

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
#
# Date:      Tue Jun 4 17:22:45 2019 -0400
#
# On branch master
# Your branch is up-to-date with 'origin/master'.
#¬
# Changes to be committed:
#      new file:   .gitignore¬
#      modified:   actions/remove.go¬
#      modified:   notes.md¬
#      modified:   readme.md¬
#
# Changes not staged for commit:
#     modified:   readme.md
```
