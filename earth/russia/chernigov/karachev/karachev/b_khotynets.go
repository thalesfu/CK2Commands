package karachev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍特涅茨KhotynetsBarony struct {
	feud.BaseBarony
}

var BKhotynets霍特涅茨 feud.Barony = &霍特涅茨KhotynetsBarony{}

func init() {
    f := BKhotynets霍特涅茨.(*霍特涅茨KhotynetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khotynets",
		TitleName: "霍特涅茨",
		TitleCode: "b_khotynets",
	}
}
