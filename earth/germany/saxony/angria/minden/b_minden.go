package minden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明登MindenBarony struct {
	feud.BaseBarony
}

var BMinden明登 feud.Barony = &明登MindenBarony{}

func init() {
	f := BMinden明登.(*明登MindenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minden",
		TitleName: "明登",
		TitleCode: "b_minden",
	}
}
