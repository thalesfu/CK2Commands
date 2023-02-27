package savolaks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉海林纳BrahelinnaBarony struct {
	feud.BaseBarony
}

var BBrahelinna布拉海林纳 feud.Barony = &布拉海林纳BrahelinnaBarony{}

func init() {
    f := BBrahelinna布拉海林纳.(*布拉海林纳BrahelinnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brahelinna",
		TitleName: "布拉海林纳",
		TitleCode: "b_brahelinna",
	}
}
