package rusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列吉诺PereginoBarony struct {
	feud.BaseBarony
}

var BPeregino佩列吉诺 feud.Barony = &佩列吉诺PereginoBarony{}

func init() {
	f := BPeregino佩列吉诺.(*佩列吉诺PereginoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peregino",
		TitleName: "佩列吉诺",
		TitleCode: "b_peregino",
	}
}
