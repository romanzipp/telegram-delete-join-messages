# bot for deleting leave/join messages in telegram

## Docker Compose Deployment

```yaml
services:
  telegram-bot:
    image: ghcr.io/romanzipp/telegram-delete-join-messages:latest
    container_name: telegram-delete-join-messages
    restart: unless-stopped
    volumes:
      - ./data:/data
    environment:
      TELEGRAM_TOKEN: "your_bot_token_here"
      TELEGRAM_ADMIN_IDS: "123456789,987654321"
      DELETE_JOIN: "true"
      DELETE_LEAVE: "true"
      RESTRICT_ON_JOIN: "false"
      RESTRICT_ON_JOIN_TIME: "600"
      ALLOWED_CHAT_IDS: ""
      INVITE_LINK: ""
      YANDEX_TOKEN: ""
      CONVERSATIONS: '[]'
      DB_PATH: "/data/bot.db"
      DEBUG: "false"
```
