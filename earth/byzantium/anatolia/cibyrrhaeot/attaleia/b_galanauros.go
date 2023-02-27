package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾拉纳罗斯GalanaurosBarony struct {
	feud.BaseBarony
}

var BGalanauros贾拉纳罗斯 feud.Barony = &贾拉纳罗斯GalanaurosBarony{}

func init() {
    f := BGalanauros贾拉纳罗斯.(*贾拉纳罗斯GalanaurosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galanauros",
		TitleName: "贾拉纳罗斯",
		TitleCode: "b_galanauros",
	}
}
