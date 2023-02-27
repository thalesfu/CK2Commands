package gudbrandsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑达尔HeidalBarony struct {
	feud.BaseBarony
}

var BHeidal黑达尔 feud.Barony = &黑达尔HeidalBarony{}

func init() {
    f := BHeidal黑达尔.(*黑达尔HeidalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heidal",
		TitleName: "黑达尔",
		TitleCode: "b_heidal",
	}
}
