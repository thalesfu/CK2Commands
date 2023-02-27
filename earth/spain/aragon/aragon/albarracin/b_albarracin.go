package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔瓦拉辛AlbarracinBarony struct {
	feud.BaseBarony
}

var BAlbarracin阿尔瓦拉辛 feud.Barony = &阿尔瓦拉辛AlbarracinBarony{}

func init() {
    f := BAlbarracin阿尔瓦拉辛.(*阿尔瓦拉辛AlbarracinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albarracin",
		TitleName: "阿尔瓦拉辛",
		TitleCode: "b_albarracin",
	}
}
