package ranthambore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 森陀摩罗尼SendamaraniBarony struct {
	feud.BaseBarony
}

var BSendamarani森陀摩罗尼 feud.Barony = &森陀摩罗尼SendamaraniBarony{}

func init() {
	f := BSendamarani森陀摩罗尼.(*森陀摩罗尼SendamaraniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sendamarani",
		TitleName: "森陀摩罗尼",
		TitleCode: "b_sendamarani",
	}
}
