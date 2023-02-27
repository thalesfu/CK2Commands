package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基亚文纳ChiavennaBarony struct {
	feud.BaseBarony
}

var BChiavenna基亚文纳 feud.Barony = &基亚文纳ChiavennaBarony{}

func init() {
    f := BChiavenna基亚文纳.(*基亚文纳ChiavennaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chiavenna",
		TitleName: "基亚文纳",
		TitleCode: "b_chiavenna",
	}
}
