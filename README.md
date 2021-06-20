# bitcloutscripts
collection of tools to gleam insights from a full bitclout node's data

```
bitcloutscripts $ ./bcs

  bcs posts       # print all posts
  bcs graph       # make clout.gv graph file
  bcs sqlite      # place clouts into local sqlite database
  bcs search      # search sqlite database

bitcloutscripts $ ./bcs posts --dir=/Users/aa/acopy/badgerdb
```

# setup, run locally

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

Let that run for ~24 hours or so, then:

```
mkdir ~/acopy
copy -r ~/.config/bitclout/bitclout/MAINNET/badgerdb ~/acopy
rm ~/acopy/badgerdb/*.mem
```

And now you can run:

```
./bcs posts --dir=/full/path/to/acopy/badgerdb
./bcs graph --dir=/full/path/to/acopy/badgerdb
etc
```

# Turnkey Script, Run on vultr Step 1

Using aws competitor <a href="https://www.vultr.com/?ref=8507322">Vultr</a> we made this pricing choice:

![pricing](https://i.imgur.com/vlOuX5Z.png)

There are cheaper than $40 a month servers but that's the bare minimum to have enough hard drive space and ram to get through the full download. At the time of this writing it will use about 65 GB to download everything.

If you get the $80 or $160 a month ones your processing might go faster but we were able to complete the task with the 160 GB drive one.

Once you have your price selected, create a new instance and select the image:

![image](https://i.imgur.com/fFDIP14.png)

ubuntu 18.04 is a nice stable version and this script has been tested with it:

https://github.com/andrewarrow/bitcloutscripts/blob/main/run.sh

# Step 2

ssh into your new instance as root:

```
ssh root@ip_address_of_new_vultr_server
```

(Note: you need to setup your sshkeys with vultr if you haven't already.)

Once you are in run the commands from the script and notice the sleep 86400 line.

When you get to that one, it'll wait a full 24 hours for the blockchain to get current.

# Step 3

After the long wait run the rest of the commands and you should see posts!

