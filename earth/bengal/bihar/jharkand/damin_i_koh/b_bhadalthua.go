package damin_i_koh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆荼罗土诃BhadalthuaBarony struct {
	feud.BaseBarony
}

var BBhadalthua婆荼罗土诃 feud.Barony = &婆荼罗土诃BhadalthuaBarony{}

func init() {
    f := BBhadalthua婆荼罗土诃.(*婆荼罗土诃BhadalthuaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadalthua",
		TitleName: "婆荼罗土诃",
		TitleCode: "b_bhadalthua",
	}
}
