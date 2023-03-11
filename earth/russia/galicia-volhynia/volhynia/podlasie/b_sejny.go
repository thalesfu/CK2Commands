package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞伊内SejnyBarony struct {
	feud.BaseBarony
}

var BSejny塞伊内 feud.Barony = &塞伊内SejnyBarony{}

func init() {
    f := BSejny塞伊内.(*塞伊内SejnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sejny",
		TitleName: "塞伊内",
		TitleCode: "b_sejny",
	}
}
