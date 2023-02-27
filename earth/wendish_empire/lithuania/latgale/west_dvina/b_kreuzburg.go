package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗伊茨堡KreuzburgBarony struct {
	feud.BaseBarony
}

var BKreuzburg克罗伊茨堡 feud.Barony = &克罗伊茨堡KreuzburgBarony{}

func init() {
    f := BKreuzburg克罗伊茨堡.(*克罗伊茨堡KreuzburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kreuzburg",
		TitleName: "克罗伊茨堡",
		TitleCode: "b_kreuzburg",
	}
}
