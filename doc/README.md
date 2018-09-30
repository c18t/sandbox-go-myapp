# Goアプリケーションの作り方

## リポジトリの作成
[個人開発の手順 - (/ɯ̹t͡ɕʲi/)](https://scrapbox.io/c18t/%E5%80%8B%E4%BA%BA%E9%96%8B%E7%99%BA%E3%81%AE%E6%89%8B%E9%A0%86)のリポジトリについてに倣う。

## 依存関係の管理
- [golang/dep](https://github.com/golang/dep)

開発環境セットアップが済んでたら動くはず。

```
$ dep init
```

`Gopkg.toml`, `Gopkg.lock`, `vendor/` が出来る。

```
$ dep ensure -add github.com/foo/bar github.com/baz/quux
```

依存関係の記録を追加するには `dep ensure -add`

### 依存関係の表示
```
$ dep status -dot | dot -T png | display
```

で依存関係を画像表示できる。

## テスト
- [testing - The Go Programming Language](https://golang.org/pkg/testing/)

**_test.go というファイル名でテストを作成する。

```
package main

import (
  "testing"
)

func TestFunc1(t *testing.T) {

}

func TestFunc2(t *testing.T) {

}

func TestFunc3(t *testing.T) {

}
```

上記の形式でテストを書き、`go test`を実施することでテストが実行される。

## コードのフォーマット
`go fmt` でファイルの記述をprettifyできる。(インデントがタブ[ハードタブ]になるのは強い思想でそうなっているので従う)

## インストール
`go get <package_repos>`でソースのダウンロードと`go install`が実行される。

このプロジェクトを例にすると、`go get github.com/c18t/sandbox-go-myapp/cmd/fizzbuzz_go`で `fizzbuzz_go`  コマンドがインストールされる。
