package nyingchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴宜DrakchiBarony struct {
	feud.BaseBarony
}

var BDrakchi巴宜 feud.Barony = &巴宜DrakchiBarony{}

func init() {
	f := BDrakchi巴宜.(*巴宜DrakchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drakchi",
		TitleName: "巴宜",
		TitleCode: "b_drakchi",
	}
}
