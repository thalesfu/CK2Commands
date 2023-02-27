package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳斯LouthBarony struct {
	feud.BaseBarony
}

var BLouth劳斯 feud.Barony = &劳斯LouthBarony{}

func init() {
    f := BLouth劳斯.(*劳斯LouthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "louth",
		TitleName: "劳斯",
		TitleCode: "b_louth",
	}
}
