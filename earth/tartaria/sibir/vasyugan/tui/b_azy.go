package tui

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿济AzyBarony struct {
	feud.BaseBarony
}

var BAzy阿济 feud.Barony = &阿济AzyBarony{}

func init() {
    f := BAzy阿济.(*阿济AzyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azy",
		TitleName: "阿济",
		TitleCode: "b_azy",
	}
}
