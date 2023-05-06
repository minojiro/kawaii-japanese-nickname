# Kawaii Japanese Nickname

List of kawaii nickname ideas in Japanese!

## Usage

```
npm install kawaii-japanese-nickname
```

```js
const japaneseNicknames = require('kawaii-japanese-nickname')

const nicknameIndex = Math.floor(Math.random() * japaneseNicknames)
const [japanese, english] = japaneseNicknames[nicknameIndex]
console.log(`🇯🇵 オッス！おら${japanese}！`)
console.log(`🇬🇧 My name is ${english}!`)
```

## Contribution

Contributions are welcome!
You can contribute by submitting pull requests from your fork to the upstream repository.

### 🇯🇵 日本語が得意な方へ！

かわいい日本語のニックネームのアイデアを、プルリクエストにて大大大募集しています！

1. `source.csv` に `みかん,mikan` のように、ひらがなとローマ字（基本的にはヘボン式ですが、字面がかわいければなんでも◎）をカンマ区切りで追記してください。
1. `add nickname: みかん` の形式のメッセージでコミットし、プルリクエストをお送りください。

🔰 はじめてのOSSコントリビュートにもぴったりです！プルリクエストの作り方などのGitの基本的な操作方法に関しては、 [こちらの記事](https://techtechmedia.com/how-to-fork-github/) がわかりやすくておすすめです！

素敵なアイデア、お待ちしております！ 🥰
