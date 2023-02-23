package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏拉SoulaBarony struct {
	feud.BaseBarony
}

var BSoula苏拉 feud.Barony = &苏拉SoulaBarony{}

func init() {
	f := BSoula苏拉.(*苏拉SoulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soula",
		TitleName: "苏拉",
		TitleCode: "b_soula",
	}
}
