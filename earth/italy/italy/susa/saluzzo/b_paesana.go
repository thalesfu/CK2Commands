package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕埃萨纳PaesanaBarony struct {
	feud.BaseBarony
}

var BPaesana帕埃萨纳 feud.Barony = &帕埃萨纳PaesanaBarony{}

func init() {
	f := BPaesana帕埃萨纳.(*帕埃萨纳PaesanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paesana",
		TitleName: "帕埃萨纳",
		TitleCode: "b_paesana",
	}
}
