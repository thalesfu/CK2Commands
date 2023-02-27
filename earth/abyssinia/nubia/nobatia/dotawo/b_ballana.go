package dotawo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉那BallanaBarony struct {
	feud.BaseBarony
}

var BBallana巴拉那 feud.Barony = &巴拉那BallanaBarony{}

func init() {
    f := BBallana巴拉那.(*巴拉那BallanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ballana",
		TitleName: "巴拉那",
		TitleCode: "b_ballana",
	}
}
