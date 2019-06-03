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
