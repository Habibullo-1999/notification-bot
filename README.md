# notification-bot





## Usage

Add your application configuration to your `.env` file in the root of your project:

```shell

TELEGRAM_TOKEN=""

DB_DRIVER=""
DB_USERNAME=""
DB_PASSWORD=""
DB_PROTOCOL=""
DB_ADDRESS=""
DB_PORT=""
DB_NAME=""
DB_PARAMS=""

WELCOME_TEXT="Добро пожаловать коллега 😉 !\n\nЭтот бот будет уведомить Вас о своевременном начале и своевременном освобождение митинг рума.\n\nУдачи 😊!"
WELCOME_TEXT_UNREGISTERED_USER="Добрый день коллега!\n\nПросим Вас зарегистрироваться на сайте http://meetingroom.alif.tj\n\nУдачи!"

UNKNOWN_COMMAND="Извините, неизвестная команда"

THIRTEEN_MINUTES_TO_START="Здравствуйте, до начало митинга осталось полчаса, пожалуйста не забудьте."
FIVE_MINUTES_TO_START="Осталось 5 минут до начало митинга, просим Вас подойти к митинг-руму"
START="Митинг начался, удачи!"
HALF_THE_TIME_HAS_PASSED="Прошло половина времени."
LEFT_A_LITTLE="Осталось мало времени, будьте в курсе."
LEFT_A_FIVE_MINUTES="Осталось 5 минут до завершение митинга, просим Вас во время закончить."
TIME_IS_OVER="Время вышло, просим Вас освободить митинг рум."
```