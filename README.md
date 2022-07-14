# parrot

Parrot is a simple passphrase generator based on [Diceware](https://theworld.com/~reinhold/diceware.html) (You can download a word list from that link).

The idea of parrot is to be very simple. By default, parrot generates 8 word length passphrases (~200bits entropy, which is security overkill for most users use cases). You can configure the output simply passing length (number) and a quoted character delimiter, `parrot generate [length] ["delimiter"]`, see example below:

```shell
$ parrot generate 6 "-"
Ella-Woo-Steak-Keats-Shift-Chile
```

## Init
Inside db, there is another directory (initDB) with a small program to initialize your database. At this moment, I'm using the diceware list, but you can initialize with your own set off words, just remember to create the list as documented on the diceware text.

### TODO

Ease database init and wordlist configuration.
