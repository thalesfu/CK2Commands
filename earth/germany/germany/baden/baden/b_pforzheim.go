package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普福尔茨海姆PforzheimBarony struct {
	feud.BaseBarony
}

var BPforzheim普福尔茨海姆 feud.Barony = &普福尔茨海姆PforzheimBarony{}

func init() {
    f := BPforzheim普福尔茨海姆.(*普福尔茨海姆PforzheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pforzheim",
		TitleName: "普福尔茨海姆",
		TitleCode: "b_pforzheim",
	}
}
