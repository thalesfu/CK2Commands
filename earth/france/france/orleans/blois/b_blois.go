package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布卢瓦BloisBarony struct {
	feud.BaseBarony
}

var BBlois布卢瓦 feud.Barony = &布卢瓦BloisBarony{}

func init() {
    f := BBlois布卢瓦.(*布卢瓦BloisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blois",
		TitleName: "布卢瓦",
		TitleCode: "b_blois",
	}
}
