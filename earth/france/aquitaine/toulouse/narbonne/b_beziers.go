package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝济耶BeziersBarony struct {
	feud.BaseBarony
}

var BBeziers贝济耶 feud.Barony = &贝济耶BeziersBarony{}

func init() {
	f := BBeziers贝济耶.(*贝济耶BeziersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beziers",
		TitleName: "贝济耶",
		TitleCode: "b_beziers",
	}
}
