package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝贾亚BejaijaBarony struct {
	feud.BaseBarony
}

var BBejaija贝贾亚 feud.Barony = &贝贾亚BejaijaBarony{}

func init() {
    f := BBejaija贝贾亚.(*贝贾亚BejaijaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bejaija",
		TitleName: "贝贾亚",
		TitleCode: "b_bejaija",
	}
}
