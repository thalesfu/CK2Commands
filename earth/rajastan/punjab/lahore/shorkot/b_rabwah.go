package shorkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉布瓦RabwahBarony struct {
	feud.BaseBarony
}

var BRabwah拉布瓦 feud.Barony = &拉布瓦RabwahBarony{}

func init() {
    f := BRabwah拉布瓦.(*拉布瓦RabwahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rabwah",
		TitleName: "拉布瓦",
		TitleCode: "b_rabwah",
	}
}
