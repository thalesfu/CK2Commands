package nakhshab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲拉布FirabrBarony struct {
	feud.BaseBarony
}

var BFirabr菲拉布 feud.Barony = &菲拉布FirabrBarony{}

func init() {
	f := BFirabr菲拉布.(*菲拉布FirabrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "firabr",
		TitleName: "菲拉布",
		TitleCode: "b_firabr",
	}
}
