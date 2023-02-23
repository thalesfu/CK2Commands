package bome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 次足拉CizulaBarony struct {
	feud.BaseBarony
}

var BCizula次足拉 feud.Barony = &次足拉CizulaBarony{}

func init() {
	f := BCizula次足拉.(*次足拉CizulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cizula",
		TitleName: "次足拉",
		TitleCode: "b_cizula",
	}
}
