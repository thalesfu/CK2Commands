package ancona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡梅里诺CamerinoBarony struct {
	feud.BaseBarony
}

var BCamerino卡梅里诺 feud.Barony = &卡梅里诺CamerinoBarony{}

func init() {
    f := BCamerino卡梅里诺.(*卡梅里诺CamerinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "camerino",
		TitleName: "卡梅里诺",
		TitleCode: "b_camerino",
	}
}
