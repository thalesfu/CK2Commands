package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布永BouillonBarony struct {
	feud.BaseBarony
}

var BBouillon布永 feud.Barony = &布永BouillonBarony{}

func init() {
    f := BBouillon布永.(*布永BouillonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bouillon",
		TitleName: "布永",
		TitleCode: "b_bouillon",
	}
}
