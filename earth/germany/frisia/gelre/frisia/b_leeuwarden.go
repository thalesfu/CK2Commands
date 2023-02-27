package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕伐登LeeuwardenBarony struct {
	feud.BaseBarony
}

var BLeeuwarden吕伐登 feud.Barony = &吕伐登LeeuwardenBarony{}

func init() {
    f := BLeeuwarden吕伐登.(*吕伐登LeeuwardenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leeuwarden",
		TitleName: "吕伐登",
		TitleCode: "b_leeuwarden",
	}
}
