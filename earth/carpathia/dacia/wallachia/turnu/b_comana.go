package turnu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科马纳ComanaBarony struct {
	feud.BaseBarony
}

var BComana科马纳 feud.Barony = &科马纳ComanaBarony{}

func init() {
	f := BComana科马纳.(*科马纳ComanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "comana",
		TitleName: "科马纳",
		TitleCode: "b_comana",
	}
}
