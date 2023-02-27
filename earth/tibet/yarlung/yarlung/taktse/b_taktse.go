package taktse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达孜TaktseBarony struct {
	feud.BaseBarony
}

var BTaktse达孜 feud.Barony = &达孜TaktseBarony{}

func init() {
    f := BTaktse达孜.(*达孜TaktseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taktse",
		TitleName: "达孜",
		TitleCode: "b_taktse",
	}
}
