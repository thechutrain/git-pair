git commit message after pair amends should look like this:

```
... message
Co-authored-by: brandon <brandon@email.com>
Co-authored-by: alan <alan@email.com>
```

```
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
