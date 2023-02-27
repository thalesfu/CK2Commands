package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾玛汀AlmatinBarony struct {
	feud.BaseBarony
}

var BAlmatin艾玛汀 feud.Barony = &艾玛汀AlmatinBarony{}

func init() {
    f := BAlmatin艾玛汀.(*艾玛汀AlmatinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almatin",
		TitleName: "艾玛汀",
		TitleCode: "b_almatin",
	}
}
