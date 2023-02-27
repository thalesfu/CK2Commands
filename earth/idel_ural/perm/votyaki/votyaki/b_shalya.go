package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙利亚ShalyaBarony struct {
	feud.BaseBarony
}

var BShalya沙利亚 feud.Barony = &沙利亚ShalyaBarony{}

func init() {
    f := BShalya沙利亚.(*沙利亚ShalyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shalya",
		TitleName: "沙利亚",
		TitleCode: "b_shalya",
	}
}
