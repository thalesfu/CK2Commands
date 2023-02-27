package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗利辛亨VlissingenBarony struct {
	feud.BaseBarony
}

var BVlissingen弗利辛亨 feud.Barony = &弗利辛亨VlissingenBarony{}

func init() {
    f := BVlissingen弗利辛亨.(*弗利辛亨VlissingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vlissingen",
		TitleName: "弗利辛亨",
		TitleCode: "b_vlissingen",
	}
}
