# Change Log

## v0.2.0

* Web API からエラーが返ってきた場合に err 扱いにする処理を追加
    * 400: [Web-API のリクエストの設定方法及び提供データの内容について（概要編）](https://www.houjin-bangou.nta.go.jp/pc/webapi/images/k-web-api-kinou-gaiyo.pdf) 末尾記載のエラーコードに対応する形で処理
    * 403: リクエスト上限エラーとして処理
    * 404: アプリケーション ID ベースの不具合として処理
    * 500: Web API 起因の不具合として処理

## v0.1.4

* Go の対応バージョンを v1.16 から v.1.20 以上に変更
* 依存パッケージのバージョンを更新
* GitHub Actions の設定も Go のバージョン変更に追従

## v0.1.3

* `Date` に対し構造体→XML, 構造体→JSON, JSON →構造体処理用のメソッドを追加
    * `MarshalXML`
    * `MarshalJSON`
    * `UnmarshalJSON`

## v0.1.2

* `ByNumbers` を削除し `ByNumber` に統一(7af1efb)

## v0.1.1

* 法人番号検索(`request.Number`)のバリデーションに法人番号のチェックデジット検証を追加
    * `byNumber`
    * `byNumberWithHistory`
    * `byNumbers`


## v0.1.0

* 法人番号システム Web-API v4 をサポートした 3 種類のリクエスト URL の生成
* レスポンス XML の構造体変換サポート
* 主要な Web-API 呼び出しメソッドの実装
