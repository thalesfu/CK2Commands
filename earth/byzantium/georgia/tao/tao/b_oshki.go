package tao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥什基OshkiBarony struct {
	feud.BaseBarony
}

var BOshki奥什基 feud.Barony = &奥什基OshkiBarony{}

func init() {
	f := BOshki奥什基.(*奥什基OshkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oshki",
		TitleName: "奥什基",
		TitleCode: "b_oshki",
	}
}
