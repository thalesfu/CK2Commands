package bharauli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆楼梨BharauliBarony struct {
	feud.BaseBarony
}

var BBharauli婆楼梨 feud.Barony = &婆楼梨BharauliBarony{}

func init() {
    f := BBharauli婆楼梨.(*婆楼梨BharauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bharauli",
		TitleName: "婆楼梨",
		TitleCode: "b_bharauli",
	}
}
