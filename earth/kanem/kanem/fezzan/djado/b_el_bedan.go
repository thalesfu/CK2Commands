package djado

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝丹El_bedanBarony struct {
	feud.BaseBarony
}

var BEl_bedan贝丹 feud.Barony = &贝丹El_bedanBarony{}

func init() {
    f := BEl_bedan贝丹.(*贝丹El_bedanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_bedan",
		TitleName: "贝丹",
		TitleCode: "b_el_bedan",
	}
}
