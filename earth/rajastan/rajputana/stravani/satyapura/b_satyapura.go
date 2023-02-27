package satyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑底也补罗SatyapuraBarony struct {
	feud.BaseBarony
}

var BSatyapura娑底也补罗 feud.Barony = &娑底也补罗SatyapuraBarony{}

func init() {
    f := BSatyapura娑底也补罗.(*娑底也补罗SatyapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satyapura",
		TitleName: "娑底也补罗",
		TitleCode: "b_satyapura",
	}
}
