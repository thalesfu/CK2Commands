package ashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁马纳RoumanaBarony struct {
	feud.BaseBarony
}

var BRoumana鲁马纳 feud.Barony = &鲁马纳RoumanaBarony{}

func init() {
	f := BRoumana鲁马纳.(*鲁马纳RoumanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roumana",
		TitleName: "鲁马纳",
		TitleCode: "b_roumana",
	}
}
