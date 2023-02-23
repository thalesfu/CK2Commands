package khuttal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 镬沙VakashBarony struct {
	feud.BaseBarony
}

var BVakash镬沙 feud.Barony = &镬沙VakashBarony{}

func init() {
	f := BVakash镬沙.(*镬沙VakashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vakash",
		TitleName: "镬沙",
		TitleCode: "b_vakash",
	}
}
