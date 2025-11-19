# bot for deleting leave/join messages in telegram

## Available Commands

### Admin Commands

Require user ID to be in `TELEGRAM_ADMIN_IDS` environment variable.

- `/kick` - Kick user from chat. Reply to target user's message.
- `/ban` - Ban user from chat permanently. Reply to target user's message.
- `/unban` - Unban user from chat. Reply to target user's message.
- `/mute` - Mute user for 1 year. Reply to target user's message.
- `/unmute` - Unmute user. Reply to target user's message.
- `/exit` - Stop the bot process.

### User Commands

- `/id` - Get your Telegram user ID and current chat ID.
- `/tldr` - Get article summary from URL using Yandex 300 API. Requires `YANDEX_TOKEN` environment variable.

## Configuration

### Deleting Join/Leave Messages

To automatically delete system messages when users join or leave a group:

1. Add the bot to your Telegram group
2. Give the bot "Delete messages" permission
3. Get your chat ID by sending `/id` in the group
4. Configure the following environment variables:

Required:
- `DELETE_JOIN="true"` - Enable deletion of join messages
- `DELETE_LEAVE="true"` - Enable deletion of leave messages
- `ALLOWED_CHAT_IDS="your_chat_id"` - Comma-separated list of chat IDs where the bot should work

Example: `ALLOWED_CHAT_IDS="-1001234567890,-1009876543210"`

The bot will only delete messages in groups listed in `ALLOWED_CHAT_IDS`.

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
      ALLOWED_CHAT_IDS: "-1001234567890"
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
  -e ALLOWED_CHAT_IDS="-1001234567890" \
  -e INVITE_LINK="" \
  -e YANDEX_TOKEN="" \
  -e CONVERSATIONS='[]' \
  -e DB_PATH="/data/bot.db" \
  -e DEBUG="false" \
  ghcr.io/romanzipp/telegram-delete-join-messages:latest
```
