/*
画像ファイルを閲覧するためのindex.htmlを生成するツール。

TABキーでページを切り替えるhtmlファイルが作られる。
プログラム実行時の引数で、画像の拡大率を変更できる。
*/
package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

const (
	TEMPLATE_HTML = "./template.html"
	INDEX_HTML    = "./index.html"
)

type ComicData struct {
	Scale  float64
	Comics map[string]map[string][]string
}

func main() {
	app := cli.NewApp()
	app.Name = "genidx"
	app.Usage = "漫画を読むためのHTML生成機"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "root,r",
			Value: ".",
			Usage: "再帰的に漫画を探すときのルート",
		},
		cli.Float64Flag{
			Name:  "scale,s",
			Value: 1.0,
			Usage: "画像の拡大率",
		},
	}

	app.Action = func(c *cli.Context) error {
		r := c.String("root")
		s := c.Float64("scale")

		paths := getAllPaths(r)
		comics := convertComicData(paths)
		hd := ComicData{
			Scale:  s,
			Comics: comics,
		}

		t := template.Must(template.ParseFiles(TEMPLATE_HTML))
		fp, err := os.Create(INDEX_HTML)
		defer fp.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(fp, hd)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	app.Run(os.Args)
}

// 再帰的に画像フォルダのパスを返す
func getAllPaths(p string) []string {
	var prevPath string
	paths := make([]string, 0)
	filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		lower := strings.ToLower(path)
		if strings.HasSuffix(lower, ".png") || strings.HasSuffix(lower, ".jpg") {
			// 漫画の変わり目に改行を区切り線を入れる
			prevCh, _ := filepath.Split(prevPath)
			currentCh, _ := filepath.Split(path)
			if prevCh != currentCh {
				//paths = append(paths, "CHANGED")
			}

			prevPath = path
			paths = append(paths, path)
		}
		return nil
	})
	return paths[1:]
}

func convertComicData(paths []string) map[string]map[string][]string {
	comics := map[string]map[string][]string{}
	for _, path := range paths {
		ps := strings.Split(path, string(os.PathSeparator))
		// 漫画のタイトル/チャプター/ページ番号
		title, ch := ps[0], ps[1]
		if comics[title] == nil {
			comics[title] = map[string][]string{}
		}
		comics[title][ch] = append(comics[title][ch], path)
	}
	return comics
}
