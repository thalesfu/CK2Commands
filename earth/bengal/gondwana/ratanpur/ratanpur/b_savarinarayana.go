package ratanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑伐利那罗衍那SavarinarayanaBarony struct {
	feud.BaseBarony
}

var BSavarinarayana娑伐利那罗衍那 feud.Barony = &娑伐利那罗衍那SavarinarayanaBarony{}

func init() {
    f := BSavarinarayana娑伐利那罗衍那.(*娑伐利那罗衍那SavarinarayanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "savarinarayana",
		TitleName: "娑伐利那罗衍那",
		TitleCode: "b_savarinarayana",
	}
}
