package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘迪亚GandiaBarony struct {
	feud.BaseBarony
}

var BGandia甘迪亚 feud.Barony = &甘迪亚GandiaBarony{}

func init() {
	f := BGandia甘迪亚.(*甘迪亚GandiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gandia",
		TitleName: "甘迪亚",
		TitleCode: "b_gandia",
	}
}
