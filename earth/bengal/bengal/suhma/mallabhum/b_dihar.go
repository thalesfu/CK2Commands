package mallabhum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提诃罗DiharBarony struct {
	feud.BaseBarony
}

var BDihar提诃罗 feud.Barony = &提诃罗DiharBarony{}

func init() {
	f := BDihar提诃罗.(*提诃罗DiharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dihar",
		TitleName: "提诃罗",
		TitleCode: "b_dihar",
	}
}
