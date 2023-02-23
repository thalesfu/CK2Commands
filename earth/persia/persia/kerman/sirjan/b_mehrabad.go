package sirjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅赫拉巴德MehrabadBarony struct {
	feud.BaseBarony
}

var BMehrabad梅赫拉巴德 feud.Barony = &梅赫拉巴德MehrabadBarony{}

func init() {
	f := BMehrabad梅赫拉巴德.(*梅赫拉巴德MehrabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mehrabad",
		TitleName: "梅赫拉巴德",
		TitleCode: "b_mehrabad",
	}
}
