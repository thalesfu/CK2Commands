package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗迦都TalakadBarony struct {
	feud.BaseBarony
}

var BTalakad多罗迦都 feud.Barony = &多罗迦都TalakadBarony{}

func init() {
    f := BTalakad多罗迦都.(*多罗迦都TalakadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talakad",
		TitleName: "多罗迦都",
		TitleCode: "b_talakad",
	}
}
