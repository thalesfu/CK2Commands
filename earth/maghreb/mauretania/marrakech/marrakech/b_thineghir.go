package marrakech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷吉尔ThineghirBarony struct {
	feud.BaseBarony
}

var BThineghir廷吉尔 feud.Barony = &廷吉尔ThineghirBarony{}

func init() {
	f := BThineghir廷吉尔.(*廷吉尔ThineghirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thineghir",
		TitleName: "廷吉尔",
		TitleCode: "b_thineghir",
	}
}
