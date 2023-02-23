package innsbruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因斯布鲁克InnsbruckBarony struct {
	feud.BaseBarony
}

var BInnsbruck因斯布鲁克 feud.Barony = &因斯布鲁克InnsbruckBarony{}

func init() {
	f := BInnsbruck因斯布鲁克.(*因斯布鲁克InnsbruckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "innsbruck",
		TitleName: "因斯布鲁克",
		TitleCode: "b_innsbruck",
	}
}
