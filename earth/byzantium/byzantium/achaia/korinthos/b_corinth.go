package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科林斯CorinthBarony struct {
	feud.BaseBarony
}

var BCorinth科林斯 feud.Barony = &科林斯CorinthBarony{}

func init() {
	f := BCorinth科林斯.(*科林斯CorinthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corinth",
		TitleName: "科林斯",
		TitleCode: "b_corinth",
	}
}
