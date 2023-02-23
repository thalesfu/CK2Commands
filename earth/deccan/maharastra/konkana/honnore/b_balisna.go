package honnore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆梨娑那BalisnaBarony struct {
	feud.BaseBarony
}

var BBalisna婆梨娑那 feud.Barony = &婆梨娑那BalisnaBarony{}

func init() {
	f := BBalisna婆梨娑那.(*婆梨娑那BalisnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balisna",
		TitleName: "婆梨娑那",
		TitleCode: "b_balisna",
	}
}
