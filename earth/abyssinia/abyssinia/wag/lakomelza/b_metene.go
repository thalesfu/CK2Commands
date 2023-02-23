package lakomelza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麼忒讷MeteneBarony struct {
	feud.BaseBarony
}

var BMetene麼忒讷 feud.Barony = &麼忒讷MeteneBarony{}

func init() {
	f := BMetene麼忒讷.(*麼忒讷MeteneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "metene",
		TitleName: "麼忒讷",
		TitleCode: "b_metene",
	}
}
