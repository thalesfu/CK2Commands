package ili

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙伦SharynBarony struct {
	feud.BaseBarony
}

var BSharyn沙伦 feud.Barony = &沙伦SharynBarony{}

func init() {
    f := BSharyn沙伦.(*沙伦SharynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharyn",
		TitleName: "沙伦",
		TitleCode: "b_sharyn",
	}
}
