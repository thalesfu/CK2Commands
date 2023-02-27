package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿达尔AdarBarony struct {
	feud.BaseBarony
}

var BAdar阿达尔 feud.Barony = &阿达尔AdarBarony{}

func init() {
    f := BAdar阿达尔.(*阿达尔AdarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adar",
		TitleName: "阿达尔",
		TitleCode: "b_adar",
	}
}
