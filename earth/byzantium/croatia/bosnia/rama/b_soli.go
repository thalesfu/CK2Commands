package rama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索利SoliBarony struct {
	feud.BaseBarony
}

var BSoli索利 feud.Barony = &索利SoliBarony{}

func init() {
	f := BSoli索利.(*索利SoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soli",
		TitleName: "索利",
		TitleCode: "b_soli",
	}
}
