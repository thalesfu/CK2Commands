package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯克Ask_romerikeBarony struct {
	feud.BaseBarony
}

var BAsk_romerike阿斯克 feud.Barony = &阿斯克Ask_romerikeBarony{}

func init() {
    f := BAsk_romerike阿斯克.(*阿斯克Ask_romerikeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ask_romerike",
		TitleName: "阿斯克",
		TitleCode: "b_ask_romerike",
	}
}
