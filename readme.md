# シン・ブログ更新内容をlineでお知らせ

[ダンススタジオのブログ](https://jko.hateblo.jp/) に更新があったらLineでお知らせするシステム。

## Motivation

PHP + さくらVPS で稼働中の [元システム](https://github.com/daisuke85a/blogPushLine) を GO + Aws Lambda でリプレイスしたい。


## Environment

```bash
$ brew install go@1.18
```

## How to use on local

まだ完成してない。単機能だけ動かせる感じ

- ブログ内容のスクレイピング
```bash
$ go run scrape.go
```

- LinePush 通知
TODO

