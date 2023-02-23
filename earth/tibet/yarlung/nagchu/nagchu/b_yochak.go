package nagchu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 油恰YochakBarony struct {
	feud.BaseBarony
}

var BYochak油恰 feud.Barony = &油恰YochakBarony{}

func init() {
	f := BYochak油恰.(*油恰YochakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yochak",
		TitleName: "油恰",
		TitleCode: "b_yochak",
	}
}
