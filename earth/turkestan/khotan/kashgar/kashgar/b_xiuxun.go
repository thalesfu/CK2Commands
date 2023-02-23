package kashgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 休循XiuxunBarony struct {
	feud.BaseBarony
}

var BXiuxun休循 feud.Barony = &休循XiuxunBarony{}

func init() {
	f := BXiuxun休循.(*休循XiuxunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiuxun",
		TitleName: "休循",
		TitleCode: "b_xiuxun",
	}
}
