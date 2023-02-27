package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图赖纳TuraynahBarony struct {
	feud.BaseBarony
}

var BTuraynah图赖纳 feud.Barony = &图赖纳TuraynahBarony{}

func init() {
    f := BTuraynah图赖纳.(*图赖纳TuraynahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turaynah",
		TitleName: "图赖纳",
		TitleCode: "b_turaynah",
	}
}
