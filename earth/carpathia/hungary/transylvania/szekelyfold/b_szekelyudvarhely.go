package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞凯尤德沃尔海伊SzekelyudvarhelyBarony struct {
	feud.BaseBarony
}

var BSzekelyudvarhely塞凯尤德沃尔海伊 feud.Barony = &塞凯尤德沃尔海伊SzekelyudvarhelyBarony{}

func init() {
    f := BSzekelyudvarhely塞凯尤德沃尔海伊.(*塞凯尤德沃尔海伊SzekelyudvarhelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szekelyudvarhely",
		TitleName: "塞凯尤德沃尔海伊",
		TitleCode: "b_szekelyudvarhely",
	}
}
