package builth

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯圭辛AbergwesynBarony struct {
	feud.BaseBarony
}

var BAbergwesyn阿伯圭辛 feud.Barony = &阿伯圭辛AbergwesynBarony{}

func init() {
	f := BAbergwesyn阿伯圭辛.(*阿伯圭辛AbergwesynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abergwesyn",
		TitleName: "阿伯圭辛",
		TitleCode: "b_abergwesyn",
	}
}
