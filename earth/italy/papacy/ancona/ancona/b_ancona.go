package ancona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安科纳AnconaBarony struct {
	feud.BaseBarony
}

var BAncona安科纳 feud.Barony = &安科纳AnconaBarony{}

func init() {
    f := BAncona安科纳.(*安科纳AnconaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ancona",
		TitleName: "安科纳",
		TitleCode: "b_ancona",
	}
}
