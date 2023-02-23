package uvs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 额尔逊ErzinBarony struct {
	feud.BaseBarony
}

var BErzin额尔逊 feud.Barony = &额尔逊ErzinBarony{}

func init() {
	f := BErzin额尔逊.(*额尔逊ErzinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erzin",
		TitleName: "额尔逊",
		TitleCode: "b_erzin",
	}
}
