package markam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 措瓦CuowaBarony struct {
	feud.BaseBarony
}

var BCuowa措瓦 feud.Barony = &措瓦CuowaBarony{}

func init() {
	f := BCuowa措瓦.(*措瓦CuowaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cuowa",
		TitleName: "措瓦",
		TitleCode: "b_cuowa",
	}
}
