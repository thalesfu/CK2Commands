package udine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰莫纳GemonaBarony struct {
	feud.BaseBarony
}

var BGemona杰莫纳 feud.Barony = &杰莫纳GemonaBarony{}

func init() {
	f := BGemona杰莫纳.(*杰莫纳GemonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gemona",
		TitleName: "杰莫纳",
		TitleCode: "b_gemona",
	}
}
