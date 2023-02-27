package kalyani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿萨AusaBarony struct {
	feud.BaseBarony
}

var BAusa阿萨 feud.Barony = &阿萨AusaBarony{}

func init() {
    f := BAusa阿萨.(*阿萨AusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ausa",
		TitleName: "阿萨",
		TitleCode: "b_ausa",
	}
}
