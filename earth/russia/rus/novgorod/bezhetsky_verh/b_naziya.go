package bezhetsky_verh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳济亚NaziyaBarony struct {
	feud.BaseBarony
}

var BNaziya纳济亚 feud.Barony = &纳济亚NaziyaBarony{}

func init() {
    f := BNaziya纳济亚.(*纳济亚NaziyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naziya",
		TitleName: "纳济亚",
		TitleCode: "b_naziya",
	}
}
