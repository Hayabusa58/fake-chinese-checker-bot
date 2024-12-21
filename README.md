# fake-chinese-checker-bot

## これは何
- Discord の特定のチャンネルにおいて漢字以外の文字列が投稿された際に"🤖不正入国者検知!" と投稿するだけのDiscord Botです
    - emoji, メンション, 記号は検知対象外です

## 使い方
### 準備
- 事前に適切な権限を与えた Discord Botを作成し、動作させたいサーバに招待しておく
- Token と検知してほしいチャンネルのIDをコピーしておく

### Docker
```
$ mv .env.sample .env

$ vi .env
# BOT_TOKEN と CHANNEL_ID を記入

$ docker build -t fake-chinese-check-bot .
$ docker run -d --rm --name fake-chinese-check-bot fake-chinese-check-bot 
```