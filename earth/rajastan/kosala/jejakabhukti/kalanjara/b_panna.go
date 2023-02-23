package kalanjara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般那PannaBarony struct {
	feud.BaseBarony
}

var BPanna般那 feud.Barony = &般那PannaBarony{}

func init() {
	f := BPanna般那.(*般那PannaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panna",
		TitleName: "般那",
		TitleCode: "b_panna",
	}
}
