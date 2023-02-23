package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔比亚OlbiaBarony struct {
	feud.BaseBarony
}

var BOlbia奥尔比亚 feud.Barony = &奥尔比亚OlbiaBarony{}

func init() {
	f := BOlbia奥尔比亚.(*奥尔比亚OlbiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olbia",
		TitleName: "奥尔比亚",
		TitleCode: "b_olbia",
	}
}
