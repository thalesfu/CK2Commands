package tadla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦乌马纳WawmanaBarony struct {
	feud.BaseBarony
}

var BWawmana瓦乌马纳 feud.Barony = &瓦乌马纳WawmanaBarony{}

func init() {
	f := BWawmana瓦乌马纳.(*瓦乌马纳WawmanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wawmana",
		TitleName: "瓦乌马纳",
		TitleCode: "b_wawmana",
	}
}
