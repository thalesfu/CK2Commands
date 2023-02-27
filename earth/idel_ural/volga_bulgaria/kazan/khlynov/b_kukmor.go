package khlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库克莫尔KukmorBarony struct {
	feud.BaseBarony
}

var BKukmor库克莫尔 feud.Barony = &库克莫尔KukmorBarony{}

func init() {
    f := BKukmor库克莫尔.(*库克莫尔KukmorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kukmor",
		TitleName: "库克莫尔",
		TitleCode: "b_kukmor",
	}
}
