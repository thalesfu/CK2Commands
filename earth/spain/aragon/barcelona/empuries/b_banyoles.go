package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尼奥莱斯BanyolesBarony struct {
	feud.BaseBarony
}

var BBanyoles巴尼奥莱斯 feud.Barony = &巴尼奥莱斯BanyolesBarony{}

func init() {
    f := BBanyoles巴尼奥莱斯.(*巴尼奥莱斯BanyolesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banyoles",
		TitleName: "巴尼奥莱斯",
		TitleCode: "b_banyoles",
	}
}
