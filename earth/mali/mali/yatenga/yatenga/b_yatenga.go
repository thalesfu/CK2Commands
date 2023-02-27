package yatenga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚滕加YatengaBarony struct {
	feud.BaseBarony
}

var BYatenga亚滕加 feud.Barony = &亚滕加YatengaBarony{}

func init() {
    f := BYatenga亚滕加.(*亚滕加YatengaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yatenga",
		TitleName: "亚滕加",
		TitleCode: "b_yatenga",
	}
}
