package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅勒MaereBarony struct {
	feud.BaseBarony
}

var BMaere梅勒 feud.Barony = &梅勒MaereBarony{}

func init() {
    f := BMaere梅勒.(*梅勒MaereBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maere",
		TitleName: "梅勒",
		TitleCode: "b_maere",
	}
}
