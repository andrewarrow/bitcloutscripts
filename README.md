# bitcloutscripts
collection of tools to gleam insights from a full bitclout node's data

```
bitcloutscripts $ ./bcs

  bcs posts       # print all posts
  bcs graph       # make clout.gv graph file

bitcloutscripts $ ./bcs posts --dir=/Users/aa/acopy/badgerdb
```

# setup

The first thing you'll need is a complete copy of the database.
But this is easy to get, you just run:

```
git clone https://github.com/bitclout/core.git
cd core
go build
./core run > /dev/null 2>&1 &
```

And it will start to download it. On a mac the location is:

~/Library/Application Support/bitclout/bitclout/MAINNET/badgerdb

On linux it is:

~/.config/bitclout/bitclout/MAINNET/badgerdb

Once you have let that run for ~24 hours or so:

```
mkdir ~/acopy
copy -r ~/.config/bitclout/bitclout/MAINNET/badgerdb ~/acopy
rm ~/acopy/badgerdb/*.mem
```

And now you can run:

./bcs posts --dir=/full/path/to/acopy/badgerdb
