package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨雷CizreBarony struct {
	feud.BaseBarony
}

var BCizre茨雷 feud.Barony = &茨雷CizreBarony{}

func init() {
	f := BCizre茨雷.(*茨雷CizreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cizre",
		TitleName: "茨雷",
		TitleCode: "b_cizre",
	}
}
