package midnapore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 檀檀DantanBarony struct {
	feud.BaseBarony
}

var BDantan檀檀 feud.Barony = &檀檀DantanBarony{}

func init() {
	f := BDantan檀檀.(*檀檀DantanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dantan",
		TitleName: "檀檀",
		TitleCode: "b_dantan",
	}
}
