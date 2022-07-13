# parrot

Parrot is a simple passphrase generator based on [Diceware](https://theworld.com/~reinhold/diceware.html) (You can download a word list from that link).

Right now, it will generate a 8 words passphrase delimited with "-", depending on the selected words, this passphrase will be between 198~205 bits, which is overkill for most use cases.

## Init

Inside db, there is another directory (initDB) with a small program to initialize your database. At this moment, I'm using the diceware list, but you can initialize with your own set off words, just remember to create the list as documented on the diceware text.

## TODO

Add command line options.
Ease database init and wordlist configuration.
