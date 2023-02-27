package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕森PasenBarony struct {
	feud.BaseBarony
}

var BPasen帕森 feud.Barony = &帕森PasenBarony{}

func init() {
    f := BPasen帕森.(*帕森PasenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pasen",
		TitleName: "帕森",
		TitleCode: "b_pasen",
	}
}
