package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 通厄伦TongerenBarony struct {
	feud.BaseBarony
}

var BTongeren通厄伦 feud.Barony = &通厄伦TongerenBarony{}

func init() {
    f := BTongeren通厄伦.(*通厄伦TongerenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tongeren",
		TitleName: "通厄伦",
		TitleCode: "b_tongeren",
	}
}
