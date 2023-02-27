package puri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯加AskaBarony struct {
	feud.BaseBarony
}

var BAska阿斯加 feud.Barony = &阿斯加AskaBarony{}

func init() {
    f := BAska阿斯加.(*阿斯加AskaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aska",
		TitleName: "阿斯加",
		TitleCode: "b_aska",
	}
}
