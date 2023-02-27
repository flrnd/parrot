# parrot

Parrot is a easy to use passphrase generator based on [Diceware](https://theworld.com/~reinhold/diceware.html) (You can download a word list from that link).

The idea of parrot is to be very simple. By default, parrot generates 8 word length passphrases (~200bits entropy, which is security overkill for most users use cases). You can configure the output simply passing length (number) and a quoted character delimiter, `parrot generate [length] ["delimiter"]`, see example below:

```shell
$ parrot generate 6 "-"
Ella-Woo-Steak-Keats-Shift-Chile
```

## Init

Download a word list from https://theworld.com/~reinhold/diceware.html (If you download one from this site, remember to remove the PGP signatures and leave just the words) or create your own.

Then run: `parrot init my_diceware_long_word_list.txt` to initialize the database.

