package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨基亚达SagiadaBarony struct {
	feud.BaseBarony
}

var BSagiada萨基亚达 feud.Barony = &萨基亚达SagiadaBarony{}

func init() {
	f := BSagiada萨基亚达.(*萨基亚达SagiadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sagiada",
		TitleName: "萨基亚达",
		TitleCode: "b_sagiada",
	}
}
