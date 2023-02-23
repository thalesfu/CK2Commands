package bidar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦摩他那KamthanaBarony struct {
	feud.BaseBarony
}

var BKamthana迦摩他那 feud.Barony = &迦摩他那KamthanaBarony{}

func init() {
	f := BKamthana迦摩他那.(*迦摩他那KamthanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamthana",
		TitleName: "迦摩他那",
		TitleCode: "b_kamthana",
	}
}
