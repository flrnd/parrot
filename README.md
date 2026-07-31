# parrot

Parrot is a simple command-line passphrase generator based on the Diceware method. It generates cryptographically secure passphrases by selecting random words from a Diceware word list using a cryptographically secure random number generator.

By default, Parrot generates **8-word** passphrases. Assuming the standard Diceware list of 7,776 words, each word contributes approximately **12.9 bits of entropy**, giving an 8-word passphrase about **103 bits of entropy**.

For most applications, this provides an extremely large security margin. Even 6-word Diceware passphrases (approximately 78 bits of entropy) are considered strong against offline attacks when used with modern password hashing algorithms.

You can customize both the number of words and the delimiter:

```shell
$ parrot generate 6 "-"
Ella-Woo-Steak-Keats-Shift-Chile
```

## Initialization

Before generating passphrases, initialize Parrot with a Diceware-compatible word list.

Download one from the original Diceware website or create your own list following the standard Diceware format.

```shell
parrot init my_diceware_word_list.txt
```

Each line in the word list must contain a five-digit Diceware index followed by its corresponding word.

## Security

The security of a Diceware passphrase comes from **entropy**, not complexity rules or character substitutions.

For a standard Diceware list:

|       Words | Approximate entropy |
| ----------: | ------------------: |
|           5 |            ~64 bits |
|           6 |            ~78 bits |
|           7 |            ~90 bits |
| 8 (default) |           ~103 bits |

The entropy is calculated as:

```
entropy = number_of_words × log₂(7776)
```

where 7,776 is the number of entries in the standard Diceware word list.

Unlike traditional passwords, Diceware passphrases are both easier to remember and resistant to brute-force attacks because every word is chosen independently using cryptographically secure randomness.

## Notes

The original Diceware method was created by Arnold G. Reinhold and remains one of the simplest and best understood techniques for generating high-entropy, human-memorable passphrases.

As an interesting anecdote, the original method literally used five physical six-sided dice to select each word. Parrot follows the same selection process, but replaces the physical dice with a cryptographically secure random number generator while preserving the same probability distribution.
