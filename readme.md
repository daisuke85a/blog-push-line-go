# シン・ブログ更新内容をlineでお知らせ

[ダンススタジオのブログ](https://jko.hateblo.jp/) に更新があったらLineでお知らせするシステム。

## Motivation

PHP + さくらVPS で稼働中の [元システム](https://github.com/daisuke85a/blogPushLine) を GO + Aws Lambda でリプレイスしたい。


## Environment

```bash
$ brew install go@1.18
$ cp .env.example .env
```
[LINE Developersコンソール](https://developers.line.biz/ja/docs/messaging-api/building-bot/#set-up-bot-on-line-developers-console) でボットを設定する
LINE Developersコンソール から LINE_BOT_CHANNEL_SECRET と LINE_BOT_CHANNEL_TOKEN を調べて、.env に 設定する

## How to use on local

まだ完成してない。単機能だけ動かせる感じ

- ブログ内容のスクレイピング
```bash
$ go run main.go line_notify.go scrape.go scrap 
```

- LinePush 通知
```bash
$ go run main.go line_notify.go scrape.go line 
```
