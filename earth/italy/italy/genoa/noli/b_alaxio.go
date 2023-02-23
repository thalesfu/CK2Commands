package noli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉克西奥AlaxioBarony struct {
	feud.BaseBarony
}

var BAlaxio阿拉克西奥 feud.Barony = &阿拉克西奥AlaxioBarony{}

func init() {
	f := BAlaxio阿拉克西奥.(*阿拉克西奥AlaxioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alaxio",
		TitleName: "阿拉克西奥",
		TitleCode: "b_alaxio",
	}
}
