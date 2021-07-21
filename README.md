# corp

`corp` パッケージは [法人番号システム Web-API](https://www.houjin-bangou.nta.go.jp/webapi/) を利用するためのパッケージです。

法人番号システム Web-API に対しリクエストを行い法人情報を取得します。
また内部的に利用している `request` パッケージはリクエストに必要な URL を生成します。

## 使用方法

[Go Doc](https://pkg.go.dev/github.com/fillin-inc/corp) に Example が用意してあります。
API の詳細については関連ドキュメントをご確認ください。

## Web-API 概要

法人番号システム Web-API には 3 種類の法人情報取得方法があります。

1. 法人番号と一致する法人情報を取得
1. 期間中に変更があった法人情報を取得
1. 法人名に一致する法人情報を取得

`corp` パッケージでは, それぞれの Web-API を簡単に利用できるメソッドを用意しています。

Web-API の検索条件をより活用した処理を行いたい場合はサブパッケージの 
[request](https://pkg.go.dev/github.com/fillin-inc/corp/request) を使うことで URL を生成することができます。

### 関連ドキュメント

- [Web-API(Ver.4.0)のリクエストの設定方法及び提供データの内容について](https://www.houjin-bangou.nta.go.jp/documents/k-web-api-kinou-ver4.pdf)
- [Web-API(Ver.1.1)のリクエストの設定方法及び提供データの内容について](https://www.houjin-bangou.nta.go.jp/documents/k-web-api-kinou-gaiyo.pdf)
    - Ver.4.0 のドキュメントから検索条件の詳細について参照されています。
- [都道府県コード](https://nlftp.mlit.go.jp/ksj/gml/codelist/PrefCd.html)
- [都道府県コード+市区町村コード](https://www.soumu.go.jp/denshijiti/code.html)

## 不具合等について

Issue から報告または Pull-Request を送ってください。

## LICENSE

MIT License
