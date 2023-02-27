package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴利香农BallyshannonBarony struct {
	feud.BaseBarony
}

var BBallyshannon巴利香农 feud.Barony = &巴利香农BallyshannonBarony{}

func init() {
    f := BBallyshannon巴利香农.(*巴利香农BallyshannonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ballyshannon",
		TitleName: "巴利香农",
		TitleCode: "b_ballyshannon",
	}
}
