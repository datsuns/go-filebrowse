package main

import (
	"github.com/lxn/walk"
	decl "github.com/lxn/walk/declarative"
)

import (
	"log"
	"path/filepath"
	"strings"
)

var (
	DeaultPath = "C:\\"
)

func main() {
	mw := &MyMainWindow{}

	win := decl.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "SearchBox",
		MinSize:  decl.Size{300, 400},
		Layout:   decl.VBox{},
		Children: []decl.Widget{
			decl.GroupBox{
				Layout: decl.HBox{},
				Children: []decl.Widget{
					decl.LineEdit{
						Text:     DeaultPath,
						AssignTo: &mw.searchBox,
					},
					decl.PushButton{
						Text:      "list",
						OnClicked: mw.clicked,
					},
				},
			},
			decl.ListBox{
				AssignTo:  &mw.listBox,
				Row:       5,
				OnKeyDown: mw.onKeyDown,
			},
		},
	}
	win.Create()
	mw.listBox.KeyUp().Attach(func(key walk.Key) {
		log.Printf("handler:%v\n", key)
	})
	//if _, err := (*win.AssignTo).Run(); err != nil {
	//	log.Fatal(err)
	//}
	(*win.AssignTo).Run()

}

//structにまとめることで、グローバル変数を作らない
type MyMainWindow struct {
	*walk.MainWindow
	searchBox *walk.LineEdit
	listBox   *walk.ListBox
}

func (mw *MyMainWindow) KeyEventHandler(key walk.Key) {
}

func (mw *MyMainWindow) clicked() {
	word := mw.searchBox.Text()
	flist, err := filepath.Glob(word + "/*")
	if err != nil {
		mw.listBox.SetModel(err.Error())
	}
	mw.listBox.SetModel(flist)
}

func (mw *MyMainWindow) onKeyDown(key walk.Key) {
	log.Printf("%v\n", key)
	switch key {
	case walk.KeyK:
		log.Printf("up!\n")
		//mw.listBox.KeyUp().Publish(walk.KeyUp)
	case walk.KeyJ:
		log.Printf("down!\n")
		//mw.listBox.KeyDown().Publish(walk.KeyDown)
	}

}

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
