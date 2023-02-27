package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉德克斯堡RadkersburgBarony struct {
	feud.BaseBarony
}

var BRadkersburg拉德克斯堡 feud.Barony = &拉德克斯堡RadkersburgBarony{}

func init() {
    f := BRadkersburg拉德克斯堡.(*拉德克斯堡RadkersburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radkersburg",
		TitleName: "拉德克斯堡",
		TitleCode: "b_radkersburg",
	}
}
