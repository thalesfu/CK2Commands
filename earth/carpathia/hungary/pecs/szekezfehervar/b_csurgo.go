package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘尔戈CsurgoBarony struct {
	feud.BaseBarony
}

var BCsurgo丘尔戈 feud.Barony = &丘尔戈CsurgoBarony{}

func init() {
    f := BCsurgo丘尔戈.(*丘尔戈CsurgoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "csurgo",
		TitleName: "丘尔戈",
		TitleCode: "b_csurgo",
	}
}
