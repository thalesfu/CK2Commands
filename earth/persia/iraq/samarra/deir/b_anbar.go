package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安巴尔AnbarBarony struct {
	feud.BaseBarony
}

var BAnbar安巴尔 feud.Barony = &安巴尔AnbarBarony{}

func init() {
	f := BAnbar安巴尔.(*安巴尔AnbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anbar",
		TitleName: "安巴尔",
		TitleCode: "b_anbar",
	}
}
