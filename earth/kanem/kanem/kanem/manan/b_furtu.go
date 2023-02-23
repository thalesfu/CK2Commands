package manan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富尔图FurtuBarony struct {
	feud.BaseBarony
}

var BFurtu富尔图 feud.Barony = &富尔图FurtuBarony{}

func init() {
	f := BFurtu富尔图.(*富尔图FurtuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "furtu",
		TitleName: "富尔图",
		TitleCode: "b_furtu",
	}
}
