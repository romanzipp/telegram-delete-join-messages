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

## Docker CLI

```bash
docker run -d \
  --name telegram-delete-join-messages \
  --restart unless-stopped \
  -v ./data:/data \
  -e TELEGRAM_TOKEN="your_bot_token_here" \
  -e TELEGRAM_ADMIN_IDS="123456789,987654321" \
  -e DELETE_JOIN="true" \
  -e DELETE_LEAVE="true" \
  -e RESTRICT_ON_JOIN="false" \
  -e RESTRICT_ON_JOIN_TIME="600" \
  -e ALLOWED_CHAT_IDS="" \
  -e INVITE_LINK="" \
  -e YANDEX_TOKEN="" \
  -e CONVERSATIONS='[]' \
  -e DB_PATH="/data/bot.db" \
  -e DEBUG="false" \
  ghcr.io/romanzipp/telegram-delete-join-messages:latest
```
