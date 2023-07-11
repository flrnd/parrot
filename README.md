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

## Strength

You can test passphrases strength with [zxcvbn](https://github.com/dropbox/zxcvbn) (Low-Budget Password Strength Estimation). An example of a 5 words passphrase:

```json
{
  "password": "Soya Prick Jura Efface Yi",
  "guesses": "559440010000000000000000",
  "guesses_log10": 23.7477535229952,
  "sequence": [
    {
      "pattern": "bruteforce",
      "token": "Soya P",
      "i": 0,
      "j": 5,
      "guesses": 1000000,
      "guesses_log10": 5.999999999999999
    },
    {
      "pattern": "dictionary",
      "i": 6,
      "j": 9,
      "token": "rick",
      "matched_word": "rick",
      "rank": 185,
      "dictionary_name": "male_names",
      "reversed": false,
      "l33t": false,
      "base_guesses": 185,
      "uppercase_variations": 1,
      "l33t_variations": 1,
      "guesses": 185,
      "guesses_log10": 2.2671717284030133
    },
    {
      "pattern": "bruteforce",
      "token": " Jura Ef",
      "i": 10,
      "j": 17,
      "guesses": 100000000,
      "guesses_log10": 8
    },
    {
      "pattern": "dictionary",
      "i": 18,
      "j": 21,
      "token": "face",
      "matched_word": "face",
      "rank": 252,
      "dictionary_name": "us_tv_and_film",
      "reversed": false,
      "l33t": false,
      "base_guesses": 252,
      "uppercase_variations": 1,
      "l33t_variations": 1,
      "guesses": 252,
      "guesses_log10": 2.4014005407815437
    },
    {
      "pattern": "bruteforce",
      "token": " Yi",
      "i": 22,
      "j": 24,
      "guesses": 1000,
      "guesses_log10": 2.9999999999999996
    }
  ],
  "calc_time": "0:00:00.002280",
  "crack_times_seconds": {
    "online_throttling_100_per_hour": "20139840360000001117985724.00",
    "online_no_throttling_10_per_second": "55944001000000000000000",
    "offline_slow_hashing_1e4_per_second": "55944001000000000000",
    "offline_fast_hashing_1e10_per_second": "55944001000000"
  },
  "crack_times_display": {
    "online_throttling_100_per_hour": "centuries",
    "online_no_throttling_10_per_second": "centuries",
    "offline_slow_hashing_1e4_per_second": "centuries",
    "offline_fast_hashing_1e10_per_second": "centuries"
  },
  "score": 4,
  "feedback": {
    "warning": "",
    "suggestions": []
  }
}
```

