package daman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉姆那迦RamnagarBarony struct {
	feud.BaseBarony
}

var BRamnagar拉姆那迦 feud.Barony = &拉姆那迦RamnagarBarony{}

func init() {
	f := BRamnagar拉姆那迦.(*拉姆那迦RamnagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramnagar",
		TitleName: "拉姆那迦",
		TitleCode: "b_ramnagar",
	}
}
