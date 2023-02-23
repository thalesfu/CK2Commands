package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯莫尔LismoreBarony struct {
	feud.BaseBarony
}

var BLismore利斯莫尔 feud.Barony = &利斯莫尔LismoreBarony{}

func init() {
	f := BLismore利斯莫尔.(*利斯莫尔LismoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lismore",
		TitleName: "利斯莫尔",
		TitleCode: "b_lismore",
	}
}
