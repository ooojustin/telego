## telego
Telegram Bot API Wrapper written in Go.

### Getting Started

Install the package by executing the following command:

```
$ go get github.com/ooojustin/telego
```

The `TelegramClient` and other Telegram related types are exposed via the `telegram` package.

```go
import "github.com/ooojustin/telego/pkg/telegram"
```

Initialize your client by passing your token to the `NewTelegramClient` function.

### Basic Example

Here's an example in which we invoke the Telegram API's [getMe method](https://core.telegram.org/bots/api#getme), which returns a [User](https://core.telegram.org/bots/api#user
). 

```go
package main

import (
    "github.com/ooojustin/telego/pkg/telegram"
    "github.com/ooojustin/telego/pkg/utils"
)

const TOKEN string = "<your token here>"

func main() {
    client := telegram.NewTelegramClient(TOKEN)
    
    if me, err := client.GetMe(); err == nil {
        utils.PrettyPrint(*me)
    } else {
        utils.Exitf(0, "testGetMe failed: %s", err)
    }
}
```
If your token is valid, the output will look like this:

```json
{
    "id": 0,
    "is_bot": true,
    "first_name": "justin",
    "last_name": "",
    "username": "justinsbot",
    "language_code": "",
    "is_premium": false,
    "added_to_attachment_menu": false,
    "can_join_groups": true,
    "can_read_all_group_messages": false,
    "supports_inline_queries": false
}
```

### Learn More

You can learn more about the API this package wraps here:

[Telegram Bot API Official Documentation](https://core.telegram.org/bots/api)
