package swetaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆荼契闷稚BadakhemundiBarony struct {
	feud.BaseBarony
}

var BBadakhemundi婆荼契闷稚 feud.Barony = &婆荼契闷稚BadakhemundiBarony{}

func init() {
    f := BBadakhemundi婆荼契闷稚.(*婆荼契闷稚BadakhemundiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badakhemundi",
		TitleName: "婆荼契闷稚",
		TitleCode: "b_badakhemundi",
	}
}
