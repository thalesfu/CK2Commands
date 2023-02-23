package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格罗AgraBarony struct {
	feud.BaseBarony
}

var BAgra阿格罗 feud.Barony = &阿格罗AgraBarony{}

func init() {
	f := BAgra阿格罗.(*阿格罗AgraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agra",
		TitleName: "阿格罗",
		TitleCode: "b_agra",
	}
}
