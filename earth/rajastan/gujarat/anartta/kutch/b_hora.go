package kutch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼罗HoraBarony struct {
	feud.BaseBarony
}

var BHora呼罗 feud.Barony = &呼罗HoraBarony{}

func init() {
    f := BHora呼罗.(*呼罗HoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hora",
		TitleName: "呼罗",
		TitleCode: "b_hora",
	}
}
