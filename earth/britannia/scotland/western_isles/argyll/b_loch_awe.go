package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥湖Loch_aweBarony struct {
	feud.BaseBarony
}

var BLoch_awe奥湖 feud.Barony = &奥湖Loch_aweBarony{}

func init() {
    f := BLoch_awe奥湖.(*奥湖Loch_aweBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loch_awe",
		TitleName: "奥湖",
		TitleCode: "b_loch_awe",
	}
}
