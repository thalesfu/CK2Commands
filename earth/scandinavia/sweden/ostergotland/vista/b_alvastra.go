package vista

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔瓦斯特拉AlvastraBarony struct {
	feud.BaseBarony
}

var BAlvastra阿尔瓦斯特拉 feud.Barony = &阿尔瓦斯特拉AlvastraBarony{}

func init() {
	f := BAlvastra阿尔瓦斯特拉.(*阿尔瓦斯特拉AlvastraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alvastra",
		TitleName: "阿尔瓦斯特拉",
		TitleCode: "b_alvastra",
	}
}
