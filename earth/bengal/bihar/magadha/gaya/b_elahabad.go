package gaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃拉哈巴德ElahabadBarony struct {
	feud.BaseBarony
}

var BElahabad埃拉哈巴德 feud.Barony = &埃拉哈巴德ElahabadBarony{}

func init() {
    f := BElahabad埃拉哈巴德.(*埃拉哈巴德ElahabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elahabad",
		TitleName: "埃拉哈巴德",
		TitleCode: "b_elahabad",
	}
}
