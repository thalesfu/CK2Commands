package nyland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛维萨LoviisaBarony struct {
	feud.BaseBarony
}

var BLoviisa洛维萨 feud.Barony = &洛维萨LoviisaBarony{}

func init() {
    f := BLoviisa洛维萨.(*洛维萨LoviisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loviisa",
		TitleName: "洛维萨",
		TitleCode: "b_loviisa",
	}
}
