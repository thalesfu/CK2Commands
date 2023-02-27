package targu_jiu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维希纳VisinaBarony struct {
	feud.BaseBarony
}

var BVisina维希纳 feud.Barony = &维希纳VisinaBarony{}

func init() {
    f := BVisina维希纳.(*维希纳VisinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "visina",
		TitleName: "维希纳",
		TitleCode: "b_visina",
	}
}
