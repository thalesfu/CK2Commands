package hisar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿私姞利呬AsigarhBarony struct {
	feud.BaseBarony
}

var BAsigarh阿私姞利呬 feud.Barony = &阿私姞利呬AsigarhBarony{}

func init() {
    f := BAsigarh阿私姞利呬.(*阿私姞利呬AsigarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asigarh",
		TitleName: "阿私姞利呬",
		TitleCode: "b_asigarh",
	}
}
