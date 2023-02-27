package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡扎尔基KazarkiBarony struct {
	feud.BaseBarony
}

var BKazarki卡扎尔基 feud.Barony = &卡扎尔基KazarkiBarony{}

func init() {
    f := BKazarki卡扎尔基.(*卡扎尔基KazarkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kazarki",
		TitleName: "卡扎尔基",
		TitleCode: "b_kazarki",
	}
}
