package ponthieu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布维尔AbbevilleBarony struct {
	feud.BaseBarony
}

var BAbbeville阿布维尔 feud.Barony = &阿布维尔AbbevilleBarony{}

func init() {
    f := BAbbeville阿布维尔.(*阿布维尔AbbevilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abbeville",
		TitleName: "阿布维尔",
		TitleCode: "b_abbeville",
	}
}
