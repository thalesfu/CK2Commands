package tebessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈斯基亚奈MaskiyanaBarony struct {
	feud.BaseBarony
}

var BMaskiyana迈斯基亚奈 feud.Barony = &迈斯基亚奈MaskiyanaBarony{}

func init() {
	f := BMaskiyana迈斯基亚奈.(*迈斯基亚奈MaskiyanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maskiyana",
		TitleName: "迈斯基亚奈",
		TitleCode: "b_maskiyana",
	}
}
