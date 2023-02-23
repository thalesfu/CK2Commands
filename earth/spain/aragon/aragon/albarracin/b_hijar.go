package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊哈尔HijarBarony struct {
	feud.BaseBarony
}

var BHijar伊哈尔 feud.Barony = &伊哈尔HijarBarony{}

func init() {
	f := BHijar伊哈尔.(*伊哈尔HijarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hijar",
		TitleName: "伊哈尔",
		TitleCode: "b_hijar",
	}
}
