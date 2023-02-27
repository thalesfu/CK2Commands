package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜恩湖Loch_doonBarony struct {
	feud.BaseBarony
}

var BLoch_doon杜恩湖 feud.Barony = &杜恩湖Loch_doonBarony{}

func init() {
    f := BLoch_doon杜恩湖.(*杜恩湖Loch_doonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loch_doon",
		TitleName: "杜恩湖",
		TitleCode: "b_loch_doon",
	}
}
