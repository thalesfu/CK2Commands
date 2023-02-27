package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥格斯堡AugsburgBarony struct {
	feud.BaseBarony
}

var BAugsburg奥格斯堡 feud.Barony = &奥格斯堡AugsburgBarony{}

func init() {
    f := BAugsburg奥格斯堡.(*奥格斯堡AugsburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "augsburg",
		TitleName: "奥格斯堡",
		TitleCode: "b_augsburg",
	}
}
