package qarazhyrya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷沙甘SaryshaganBarony struct {
	feud.BaseBarony
}

var BSaryshagan萨雷沙甘 feud.Barony = &萨雷沙甘SaryshaganBarony{}

func init() {
	f := BSaryshagan萨雷沙甘.(*萨雷沙甘SaryshaganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saryshagan",
		TitleName: "萨雷沙甘",
		TitleCode: "b_saryshagan",
	}
}
