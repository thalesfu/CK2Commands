package arjin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿牙克库木AyakkumBarony struct {
	feud.BaseBarony
}

var BAyakkum阿牙克库木 feud.Barony = &阿牙克库木AyakkumBarony{}

func init() {
	f := BAyakkum阿牙克库木.(*阿牙克库木AyakkumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayakkum",
		TitleName: "阿牙克库木",
		TitleCode: "b_ayakkum",
	}
}
