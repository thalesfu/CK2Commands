package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加扎勒AlbaqarBarony struct {
	feud.BaseBarony
}

var BAlbaqar加扎勒 feud.Barony = &加扎勒AlbaqarBarony{}

func init() {
    f := BAlbaqar加扎勒.(*加扎勒AlbaqarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albaqar",
		TitleName: "加扎勒",
		TitleCode: "b_albaqar",
	}
}
