package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马加斯MagasBarony struct {
	feud.BaseBarony
}

var BMagas马加斯 feud.Barony = &马加斯MagasBarony{}

func init() {
    f := BMagas马加斯.(*马加斯MagasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "magas",
		TitleName: "马加斯",
		TitleCode: "b_magas",
	}
}
