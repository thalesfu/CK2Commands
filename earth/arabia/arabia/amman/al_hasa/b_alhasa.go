package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈萨AlhasaBarony struct {
	feud.BaseBarony
}

var BAlhasa哈萨 feud.Barony = &哈萨AlhasaBarony{}

func init() {
    f := BAlhasa哈萨.(*哈萨AlhasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhasa",
		TitleName: "哈萨",
		TitleCode: "b_alhasa",
	}
}
