package ponthieu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗利克斯库尔FlixecourtBarony struct {
	feud.BaseBarony
}

var BFlixecourt弗利克斯库尔 feud.Barony = &弗利克斯库尔FlixecourtBarony{}

func init() {
    f := BFlixecourt弗利克斯库尔.(*弗利克斯库尔FlixecourtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "flixecourt",
		TitleName: "弗利克斯库尔",
		TitleCode: "b_flixecourt",
	}
}
