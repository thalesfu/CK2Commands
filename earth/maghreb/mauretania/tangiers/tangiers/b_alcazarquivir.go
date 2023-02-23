package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔卡萨基维尔AlcazarquivirBarony struct {
	feud.BaseBarony
}

var BAlcazarquivir阿尔卡萨基维尔 feud.Barony = &阿尔卡萨基维尔AlcazarquivirBarony{}

func init() {
	f := BAlcazarquivir阿尔卡萨基维尔.(*阿尔卡萨基维尔AlcazarquivirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcazarquivir",
		TitleName: "阿尔卡萨基维尔",
		TitleCode: "b_alcazarquivir",
	}
}
