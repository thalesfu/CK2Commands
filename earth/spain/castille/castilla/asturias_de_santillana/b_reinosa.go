package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷诺萨ReinosaBarony struct {
	feud.BaseBarony
}

var BReinosa雷诺萨 feud.Barony = &雷诺萨ReinosaBarony{}

func init() {
    f := BReinosa雷诺萨.(*雷诺萨ReinosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reinosa",
		TitleName: "雷诺萨",
		TitleCode: "b_reinosa",
	}
}
