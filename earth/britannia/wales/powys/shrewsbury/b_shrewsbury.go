package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什鲁斯伯里ShrewsburyBarony struct {
	feud.BaseBarony
}

var BShrewsbury什鲁斯伯里 feud.Barony = &什鲁斯伯里ShrewsburyBarony{}

func init() {
	f := BShrewsbury什鲁斯伯里.(*什鲁斯伯里ShrewsburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shrewsbury",
		TitleName: "什鲁斯伯里",
		TitleCode: "b_shrewsbury",
	}
}
