package medjerda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈杰尔达MedjerdaBarony struct {
	feud.BaseBarony
}

var BMedjerda迈杰尔达 feud.Barony = &迈杰尔达MedjerdaBarony{}

func init() {
    f := BMedjerda迈杰尔达.(*迈杰尔达MedjerdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medjerda",
		TitleName: "迈杰尔达",
		TitleCode: "b_medjerda",
	}
}
