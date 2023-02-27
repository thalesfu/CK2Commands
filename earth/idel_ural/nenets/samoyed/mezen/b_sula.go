package mezen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏拉SulaBarony struct {
	feud.BaseBarony
}

var BSula苏拉 feud.Barony = &苏拉SulaBarony{}

func init() {
    f := BSula苏拉.(*苏拉SulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sula",
		TitleName: "苏拉",
		TitleCode: "b_sula",
	}
}
