package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格哈纳ArghanaBarony struct {
	feud.BaseBarony
}

var BArghana阿格哈纳 feud.Barony = &阿格哈纳ArghanaBarony{}

func init() {
	f := BArghana阿格哈纳.(*阿格哈纳ArghanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arghana",
		TitleName: "阿格哈纳",
		TitleCode: "b_arghana",
	}
}
