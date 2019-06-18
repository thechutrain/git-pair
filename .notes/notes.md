git commit message after pair amends should look like this:

```
... message

Co-authored-by: brandon <brandon@email.com>
Co-authored-by: alan <alan@email.com>
# git comments
#
```

```sh
#!/bin/sh
set -e

CURR_PAIR_FILE=$HOME/.pear/$(basename $PWD)/current_pairs.txt
cat $CURR_PAIR_FILE | awk 'BEGIN{FS=" "} {print "Co-authored-by: " $1 " <" $2 ">"}' >> $1
```

prepare-commit-msg hook:

```sh
#!/bin/sh
set -e

CURR_PAIR_FILE=$HOME/.pear/$(basename $PWD)/current_pairs.txt
TEMP="$HOME/.pear/.default-commit"

# Save the old content.
cat $1 > $TEMP

# Prepend two newlines.
printf "\n\n" > $1

# Create the Co-authored-by tags.
cat $CURR_PAIR_FILE | awk 'BEGIN{FS=" "} {print "Co-authored-by: " $1 " <" $2 ">"}' >> $1

# Append the old content.
cat $TEMP | sed "/^Co-authored-by/d" >> $1

rm $TEMP
```

## Saving current pairs

- you can use the `.git/config` file
- use `git config` commands to modify these commands:
  - http://craig-russell.co.uk/2011/08/24/git-tip-custom-config-parameters.html#.XPaUBtNKiV4

```
// Note throws an error that "key does not contain a section" if you try to get something that doesnt exist!


git config --add pair.coauthor namehere // you need the --add flag in order to not overwrite prev val
git config --get-all pair.coauthor // gets all the coauthors
git config --unset pair.coauthor nameofpair // removes a single user
git config --unset-all pair.coauthor // removes all users
```

## Misc reading arguments from a shell script

```sh
#!/bin/sh

if [ -z $1 ]; then
    echo empty
else
    echo not empty
fi
```

## Set up

Add this to your `.git/hooks/prepare-commit-msg` file

```sh
#!/bin/sh

set -e

# cannot alias this?
gitpair _prepare-commit-msg $@ #adds all of the arguments in bash
```
