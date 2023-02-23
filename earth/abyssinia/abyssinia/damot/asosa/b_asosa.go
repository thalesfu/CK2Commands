package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿索萨AsosaBarony struct {
	feud.BaseBarony
}

var BAsosa阿索萨 feud.Barony = &阿索萨AsosaBarony{}

func init() {
	f := BAsosa阿索萨.(*阿索萨AsosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asosa",
		TitleName: "阿索萨",
		TitleCode: "b_asosa",
	}
}
