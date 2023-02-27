package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗福德LongfordBarony struct {
	feud.BaseBarony
}

var BLongford朗福德 feud.Barony = &朗福德LongfordBarony{}

func init() {
    f := BLongford朗福德.(*朗福德LongfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "longford",
		TitleName: "朗福德",
		TitleCode: "b_longford",
	}
}
