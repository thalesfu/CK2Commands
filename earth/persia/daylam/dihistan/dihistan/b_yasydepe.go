package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚瑟杰佩YasydepeBarony struct {
	feud.BaseBarony
}

var BYasydepe亚瑟杰佩 feud.Barony = &亚瑟杰佩YasydepeBarony{}

func init() {
    f := BYasydepe亚瑟杰佩.(*亚瑟杰佩YasydepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yasydepe",
		TitleName: "亚瑟杰佩",
		TitleCode: "b_yasydepe",
	}
}
