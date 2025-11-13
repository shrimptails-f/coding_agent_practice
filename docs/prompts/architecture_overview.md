# アーキテクチャ概要（Sample Edition）

このドキュメントは、`sample_controller` を起点とした最小クリーンアーキテクチャ構成を説明します。複雑な外部依存を排除しつつ、各レイヤーの責務やデータフローが一目で分かるようにしています。

---

## 1. レイヤー構成
| レイヤー | ディレクトリ | 役割 |
| --- | --- | --- |
| Presentation | `internal/app/presentation/sample_controller.go` | Gin ハンドラ。JSON の `title` を受け取り、201で空ボディを返す。 |
| Router/Server | `internal/app/router`, `internal/app/server` | ルート定義と HTTP サーバ起動。DI コンテナからコントローラを取り出してエンドポイントへ割り当て。 |
| Application | `internal/sample/application` | ユースケース。ドメインルールを組み立て、Repository/IDProvider/Clock とやり取りする。 |
| Domain | `internal/sample/domain` | エンティティと入力検証 (`CreateSampleInput`) を保持。副作用ゼロの純粋なコードだけ。 |
| Infrastructure | `internal/sample/infrastructure` | GORM Repository（`samplese` テーブルに保存）/ UUID Provider / System Clock など。 |
| DI | `internal/di` | go.uber.org/dig で依存解決。各レイヤーを疎結合のまま束ねる。 |

---

## 2. データフロー（`POST /samples`）
```
curl → Gin Router → sample_controller → sample.UseCase → MemoryRepository
```
1. `sample_controller.CreateSample` が JSON (`title` のみ) を `CreateSampleInput` にマッピング。
2. アプリケーション層で `Validate()` を実行し、UUID と現在時刻を注入。
3. `GormRepository.Save` が GORM の `Create` によって `samplese` テーブルへ保存。
4. 成功時はボディなしで `201 Created` を返却。

`GET /samples` も同様に、アプリケーション層から Repository を経由してデータを取得します。

---

## 3. 依存性注入
- `internal/di/sample.go` で Repository / UUIDProvider / Clock を Provide。
- `internal/di/presentation.go` で UseCase を受け取り `SampleController` を構築。
- `internal/di/container.go` が `BuildContainer()` を公開し、`server.Run()` が呼び出す。

DI によって、コントローラやユースケースのテスト時は任意のモック実装に差し替え可能です。`internal/sample/application/usecase_test.go` では、インターフェースを満たすスタブを使ってユースケース単独の検証を行っています。

---

## 4. 拡張のヒント
1. 別 DB を使いたい場合は `internal/sample/infrastructure` に Repository 実装を追加し、DI の Provide 先を差し替える。
2. バリデーションルールを増やす場合は `domain` に集約し、アプリ層の責務を薄く保つ。
3. 新しいエンドポイントは `presentation` に薄いハンドラを足し、UseCase を介してドメインに渡す。

このサンプルをベースに、必要な依存やレイヤーを段階的に足していけば、本番環境向けの構成へスムーズに移行できます。
