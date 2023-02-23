package munda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那兜NathorlBarony struct {
	feud.BaseBarony
}

var BNathorl那兜 feud.Barony = &那兜NathorlBarony{}

func init() {
	f := BNathorl那兜.(*那兜NathorlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nathorl",
		TitleName: "那兜",
		TitleCode: "b_nathorl",
	}
}
