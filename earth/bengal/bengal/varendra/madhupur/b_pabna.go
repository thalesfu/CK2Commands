package madhupur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波婆那PabnaBarony struct {
	feud.BaseBarony
}

var BPabna波婆那 feud.Barony = &波婆那PabnaBarony{}

func init() {
	f := BPabna波婆那.(*波婆那PabnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pabna",
		TitleName: "波婆那",
		TitleCode: "b_pabna",
	}
}
