package jaen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞古拉SeguraBarony struct {
	feud.BaseBarony
}

var BSegura塞古拉 feud.Barony = &塞古拉SeguraBarony{}

func init() {
	f := BSegura塞古拉.(*塞古拉SeguraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "segura",
		TitleName: "塞古拉",
		TitleCode: "b_segura",
	}
}
