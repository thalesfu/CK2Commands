package tadla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达伊DaiBarony struct {
	feud.BaseBarony
}

var BDai达伊 feud.Barony = &达伊DaiBarony{}

func init() {
    f := BDai达伊.(*达伊DaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dai",
		TitleName: "达伊",
		TitleCode: "b_dai",
	}
}
