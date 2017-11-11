# genidx
フォルダ配下に存在する漫画を読むためHTML生成ツール

# 使い方

同階層に漫画などのフォルダが存在する状態で下記のコマンドを実行する。
```
./genidx
```

## ヘルプメッセージを表示する

```
./genidx h
```

```
~/Pictures/comic$ ./genidx h
NAME:
   genidx - 漫画を読むためのHTML生成機

USAGE:
   genidx [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --root value, -r value   再帰的に漫画を探すときのルート (default: ".")
   --scale value, -s value  画像の拡大率 (default: 1)
   --help, -h               show help
   --version, -v            print the version
```
