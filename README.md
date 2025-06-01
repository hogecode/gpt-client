# gpt-client


OpenAIAPIに質問し、markdownやhtmlで回答を保存できるCLIアプリです。まだほとんど開発していないので動作しません。

個人的に書籍などを読んでいて、chatgptに質問したいが回答に時間がかかるので、非同期で呼び出し、後で回答を確認したいと思ったのが制作のきっかけです。

### コマンド一覧

#### 設定関連コマンド

- `gptcl config --api-key <KEY>`
    - APIキーをJSONに保存し、動作確認
- `gptcl config --output-dir ~/gpt_outputs`
    - 出力ファイル（Markdown/HTML）の保存先ディレクトリを指定
- `gptcl config --show`
    - 現在の設定内容を表示
- `gptcl check`
    - APIキーが正常かどうか、問い合わせで接続確認

#### クエリ送信コマンド

- `gptcl --query "こんにちは" (--format markdown | html)` 
    -  クエリを送信し、成功かどうかのみ表示
- `gptcl --file queries.txt (--format markdown | html)`
    - 複数行のクエリをファイルから読み込んで連続送信
   
#### 対話履歴

- `gptcl history`
    - SQLiteに保存されたすべての対話履歴を一覧表示
- `gptcl history --fromId 5 (--toId 10) `
    - ページネーション付きで履歴を表示

#### エキスポート

- `gptcl export --format markdown | html`
    - すべての対話履歴を Markdown, HTML 形式で出力
- `gptcl export --fromId 5 (--toId 10)  --format markdown | html `
    - ページネーション付きでエキスポート

#### ディレクトリ構造

```plaintext
~/.gptcl/
├── config.json
├── chat_history.db  
```

### 仕様技術
