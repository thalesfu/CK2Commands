package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔法EdesurfaBarony struct {
	feud.BaseBarony
}

var BEdesurfa乌尔法 feud.Barony = &乌尔法EdesurfaBarony{}

func init() {
	f := BEdesurfa乌尔法.(*乌尔法EdesurfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "edesurfa",
		TitleName: "乌尔法",
		TitleCode: "b_edesurfa",
	}
}
