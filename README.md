# SuperMemo 2 telegram bot (monorepo)

Description: https://en.wikipedia.org/wiki/SuperMemo

Services:

* [Infrastructure](infrastructure) - Development infrastructure environment
* [Card](card) - Card service
* [Telegram bot](telegram-bot) - Service for interacting with Telegram
* [Telegram gate](telegram-gate) - Service for sending messages to Telegram

Reference:
* SM2 (anki) core: https://github.com/open-spaced-repetition/anki-sm-2

## Known issues

For kafka_data folder need change permissions 1001:1001
```chown 1001:1001 kafka_data```