package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰拉格LetelaghBarony struct {
	feud.BaseBarony
}

var BLetelagh泰拉格 feud.Barony = &泰拉格LetelaghBarony{}

func init() {
    f := BLetelagh泰拉格.(*泰拉格LetelaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "letelagh",
		TitleName: "泰拉格",
		TitleCode: "b_letelagh",
	}
}
