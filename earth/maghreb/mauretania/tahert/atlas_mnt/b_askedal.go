package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯克达尔AskedalBarony struct {
	feud.BaseBarony
}

var BAskedal阿斯克达尔 feud.Barony = &阿斯克达尔AskedalBarony{}

func init() {
    f := BAskedal阿斯克达尔.(*阿斯克达尔AskedalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askedal",
		TitleName: "阿斯克达尔",
		TitleCode: "b_askedal",
	}
}
