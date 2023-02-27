package eastern_sayan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨彦Eastern_sayanBarony struct {
	feud.BaseBarony
}

var BEastern_sayan萨彦 feud.Barony = &萨彦Eastern_sayanBarony{}

func init() {
    f := BEastern_sayan萨彦.(*萨彦Eastern_sayanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eastern_sayan",
		TitleName: "萨彦",
		TitleCode: "b_eastern_sayan",
	}
}
