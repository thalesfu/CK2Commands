package paro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 虎穴TaktshangBarony struct {
	feud.BaseBarony
}

var BTaktshang虎穴 feud.Barony = &虎穴TaktshangBarony{}

func init() {
	f := BTaktshang虎穴.(*虎穴TaktshangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taktshang",
		TitleName: "虎穴",
		TitleCode: "b_taktshang",
	}
}
