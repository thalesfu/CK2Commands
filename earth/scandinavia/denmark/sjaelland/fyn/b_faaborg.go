package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福堡FaaborgBarony struct {
	feud.BaseBarony
}

var BFaaborg福堡 feud.Barony = &福堡FaaborgBarony{}

func init() {
	f := BFaaborg福堡.(*福堡FaaborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faaborg",
		TitleName: "福堡",
		TitleCode: "b_faaborg",
	}
}
