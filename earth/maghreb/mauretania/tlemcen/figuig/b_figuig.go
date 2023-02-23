package figuig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲吉格FiguigBarony struct {
	feud.BaseBarony
}

var BFiguig菲吉格 feud.Barony = &菲吉格FiguigBarony{}

func init() {
	f := BFiguig菲吉格.(*菲吉格FiguigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "figuig",
		TitleName: "菲吉格",
		TitleCode: "b_figuig",
	}
}
