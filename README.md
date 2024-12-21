# fake-chinese-checker-bot

## これは何
- Discord の特定のチャンネルにおいて偽中国語以外の文字列(英数字|ひらがな|かたかな)が投稿された際に"🤖不正入国者検知!" と検知するだけのDiscord Botです
- emoji 単体のメッセージは検知しません

## 使い方
### 準備
- 事前に適切な権限を与えた Discord Botを作成し、動作させたいサーバに招待しておく
- Token と検知してほしいチャンネルのIDをメモっておく

### Docker
```
$ mv .env.sample .env

$ vi .env
# BOT_TOKEN と CHANNEL_ID を記入

$ docker build -t fake-chinese-check-bot .
$ docker run -d --rm --name fake-chinese-check-bot fake-chinese-check-bot 
```

## ToDo
- emoji と偽中国語が混合したメッセージの場合、漢字以外が含まれていなくても検知してしまう