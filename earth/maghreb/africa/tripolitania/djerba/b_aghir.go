package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾吉尔AghirBarony struct {
	feud.BaseBarony
}

var BAghir艾吉尔 feud.Barony = &艾吉尔AghirBarony{}

func init() {
    f := BAghir艾吉尔.(*艾吉尔AghirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aghir",
		TitleName: "艾吉尔",
		TitleCode: "b_aghir",
	}
}
