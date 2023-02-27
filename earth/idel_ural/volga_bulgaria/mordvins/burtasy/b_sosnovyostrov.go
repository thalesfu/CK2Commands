package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松岛SosnovyostrovBarony struct {
	feud.BaseBarony
}

var BSosnovyostrov松岛 feud.Barony = &松岛SosnovyostrovBarony{}

func init() {
    f := BSosnovyostrov松岛.(*松岛SosnovyostrovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sosnovyostrov",
		TitleName: "松岛",
		TitleCode: "b_sosnovyostrov",
	}
}
