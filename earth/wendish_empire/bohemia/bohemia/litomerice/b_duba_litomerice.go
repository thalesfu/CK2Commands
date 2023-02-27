package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜巴Duba_litomericeBarony struct {
	feud.BaseBarony
}

var BDuba_litomerice杜巴 feud.Barony = &杜巴Duba_litomericeBarony{}

func init() {
    f := BDuba_litomerice杜巴.(*杜巴Duba_litomericeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duba_litomerice",
		TitleName: "杜巴",
		TitleCode: "b_duba_litomerice",
	}
}
