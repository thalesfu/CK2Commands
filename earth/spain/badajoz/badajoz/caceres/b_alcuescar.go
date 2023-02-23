package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔库埃斯卡尔AlcuescarBarony struct {
	feud.BaseBarony
}

var BAlcuescar阿尔库埃斯卡尔 feud.Barony = &阿尔库埃斯卡尔AlcuescarBarony{}

func init() {
	f := BAlcuescar阿尔库埃斯卡尔.(*阿尔库埃斯卡尔AlcuescarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcuescar",
		TitleName: "阿尔库埃斯卡尔",
		TitleCode: "b_alcuescar",
	}
}
