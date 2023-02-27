package assab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼达MandaBarony struct {
	feud.BaseBarony
}

var BManda曼达 feud.Barony = &曼达MandaBarony{}

func init() {
    f := BManda曼达.(*曼达MandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manda",
		TitleName: "曼达",
		TitleCode: "b_manda",
	}
}
