package main

//WALK関連のライブラリ
import (
	"github.com/lxn/walk"
	decl "github.com/lxn/walk/declarative"
)

//その他標準ライブラリ
import (
	"log"
	"path/filepath"
	"strings"
)

func main() {
	mw := &MyMainWindow{}

	if _, err := (decl.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "SearchBox",
		MinSize:  decl.Size{300, 400},
		Layout:   decl.VBox{},
		Children: []decl.Widget{
			decl.GroupBox{
				Layout: decl.HBox{},
				Children: []decl.Widget{
					decl.LineEdit{
						AssignTo: &mw.searchBox,
					},
					decl.PushButton{
						Text:      "list",
						OnClicked: mw.clicked,
					},
				},
			},
			decl.ListBox{
				AssignTo: &mw.results,
				Row:      5,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

}

//structにまとめることで、グローバル変数を作らない
type MyMainWindow struct {
	*walk.MainWindow
	searchBox *walk.LineEdit
	results   *walk.ListBox
}

func (mw *MyMainWindow) clicked() {
	word := mw.searchBox.Text()
	flist, err := filepath.Glob(word + "/*")
	if err != nil {
		mw.results.SetModel(err.Error())
	}
	mw.results.SetModel(flist)
}

//textからwordを検索して、位置をUnicode単位で返す
func search(text, word string) (result []int) {
	result = []int{}
	i := 0
	for j, _ := range text {
		if strings.HasPrefix(text[j:], word) {
			log.Print(i)
			result = append(result, i)
		}
		i += 1
	}
	return
}
