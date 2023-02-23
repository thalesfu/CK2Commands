package tver

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克林KlinBarony struct {
	feud.BaseBarony
}

var BKlin克林 feud.Barony = &克林KlinBarony{}

func init() {
	f := BKlin克林.(*克林KlinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klin",
		TitleName: "克林",
		TitleCode: "b_klin",
	}
}
