package napata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加扎利GhazaliBarony struct {
	feud.BaseBarony
}

var BGhazali加扎利 feud.Barony = &加扎利GhazaliBarony{}

func init() {
	f := BGhazali加扎利.(*加扎利GhazaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghazali",
		TitleName: "加扎利",
		TitleCode: "b_ghazali",
	}
}
