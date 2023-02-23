package tripolitana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿贾米拉AljamilBarony struct {
	feud.BaseBarony
}

var BAljamil阿贾米拉 feud.Barony = &阿贾米拉AljamilBarony{}

func init() {
	f := BAljamil阿贾米拉.(*阿贾米拉AljamilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljamil",
		TitleName: "阿贾米拉",
		TitleCode: "b_aljamil",
	}
}
