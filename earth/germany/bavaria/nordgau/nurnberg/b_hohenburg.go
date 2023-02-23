package nurnberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍恩堡HohenburgBarony struct {
	feud.BaseBarony
}

var BHohenburg霍恩堡 feud.Barony = &霍恩堡HohenburgBarony{}

func init() {
	f := BHohenburg霍恩堡.(*霍恩堡HohenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hohenburg",
		TitleName: "霍恩堡",
		TitleCode: "b_hohenburg",
	}
}
