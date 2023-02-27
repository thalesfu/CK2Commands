package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尼西堡EnisiburgBarony struct {
	feud.BaseBarony
}

var BEnisiburg埃尼西堡 feud.Barony = &埃尼西堡EnisiburgBarony{}

func init() {
    f := BEnisiburg埃尼西堡.(*埃尼西堡EnisiburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "enisiburg",
		TitleName: "埃尼西堡",
		TitleCode: "b_enisiburg",
	}
}
