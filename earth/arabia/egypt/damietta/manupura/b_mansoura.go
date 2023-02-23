package manupura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼苏雷MansouraBarony struct {
	feud.BaseBarony
}

var BMansoura曼苏雷 feud.Barony = &曼苏雷MansouraBarony{}

func init() {
	f := BMansoura曼苏雷.(*曼苏雷MansouraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mansoura",
		TitleName: "曼苏雷",
		TitleCode: "b_mansoura",
	}
}
