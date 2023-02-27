package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳欧扎LalaouzaBarony struct {
	feud.BaseBarony
}

var BLalaouza劳欧扎 feud.Barony = &劳欧扎LalaouzaBarony{}

func init() {
    f := BLalaouza劳欧扎.(*劳欧扎LalaouzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lalaouza",
		TitleName: "劳欧扎",
		TitleCode: "b_lalaouza",
	}
}
