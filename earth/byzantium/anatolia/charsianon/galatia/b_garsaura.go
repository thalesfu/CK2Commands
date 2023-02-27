package galatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加苏拉GarsauraBarony struct {
	feud.BaseBarony
}

var BGarsaura加苏拉 feud.Barony = &加苏拉GarsauraBarony{}

func init() {
    f := BGarsaura加苏拉.(*加苏拉GarsauraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garsaura",
		TitleName: "加苏拉",
		TitleCode: "b_garsaura",
	}
}
