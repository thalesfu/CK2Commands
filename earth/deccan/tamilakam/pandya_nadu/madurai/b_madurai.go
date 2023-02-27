package madurai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩菟来MaduraiBarony struct {
	feud.BaseBarony
}

var BMadurai摩菟来 feud.Barony = &摩菟来MaduraiBarony{}

func init() {
    f := BMadurai摩菟来.(*摩菟来MaduraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madurai",
		TitleName: "摩菟来",
		TitleCode: "b_madurai",
	}
}
