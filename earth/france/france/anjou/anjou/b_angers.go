package anjou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂热AngersBarony struct {
	feud.BaseBarony
}

var BAngers昂热 feud.Barony = &昂热AngersBarony{}

func init() {
	f := BAngers昂热.(*昂热AngersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "angers",
		TitleName: "昂热",
		TitleCode: "b_angers",
	}
}
