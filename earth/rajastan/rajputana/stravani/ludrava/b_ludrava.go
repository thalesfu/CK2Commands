package ludrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 律陀罗婆LudravaBarony struct {
	feud.BaseBarony
}

var BLudrava律陀罗婆 feud.Barony = &律陀罗婆LudravaBarony{}

func init() {
	f := BLudrava律陀罗婆.(*律陀罗婆LudravaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ludrava",
		TitleName: "律陀罗婆",
		TitleCode: "b_ludrava",
	}
}
