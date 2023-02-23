package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索舍勒SorseleBarony struct {
	feud.BaseBarony
}

var BSorsele索舍勒 feud.Barony = &索舍勒SorseleBarony{}

func init() {
	f := BSorsele索舍勒.(*索舍勒SorseleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorsele",
		TitleName: "索舍勒",
		TitleCode: "b_sorsele",
	}
}
