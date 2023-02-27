package coruna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿科鲁尼亚CorunnaBarony struct {
	feud.BaseBarony
}

var BCorunna阿科鲁尼亚 feud.Barony = &阿科鲁尼亚CorunnaBarony{}

func init() {
    f := BCorunna阿科鲁尼亚.(*阿科鲁尼亚CorunnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corunna",
		TitleName: "阿科鲁尼亚",
		TitleCode: "b_corunna",
	}
}
