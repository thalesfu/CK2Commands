package gwynedd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯古因格雷金AbergwyngregynBarony struct {
	feud.BaseBarony
}

var BAbergwyngregyn阿伯古因格雷金 feud.Barony = &阿伯古因格雷金AbergwyngregynBarony{}

func init() {
	f := BAbergwyngregyn阿伯古因格雷金.(*阿伯古因格雷金AbergwyngregynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abergwyngregyn",
		TitleName: "阿伯古因格雷金",
		TitleCode: "b_abergwyngregyn",
	}
}
