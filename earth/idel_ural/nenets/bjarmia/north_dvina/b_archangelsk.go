package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔汉格尔斯克ArchangelskBarony struct {
	feud.BaseBarony
}

var BArchangelsk阿尔汉格尔斯克 feud.Barony = &阿尔汉格尔斯克ArchangelskBarony{}

func init() {
    f := BArchangelsk阿尔汉格尔斯克.(*阿尔汉格尔斯克ArchangelskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "archangelsk",
		TitleName: "阿尔汉格尔斯克",
		TitleCode: "b_archangelsk",
	}
}
