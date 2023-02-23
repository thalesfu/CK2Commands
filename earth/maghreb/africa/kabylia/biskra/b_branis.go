package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卜拉尼斯BranisBarony struct {
	feud.BaseBarony
}

var BBranis卜拉尼斯 feud.Barony = &卜拉尼斯BranisBarony{}

func init() {
	f := BBranis卜拉尼斯.(*卜拉尼斯BranisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "branis",
		TitleName: "卜拉尼斯",
		TitleCode: "b_branis",
	}
}
