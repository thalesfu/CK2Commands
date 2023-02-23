package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯海尔托亨博思ShertogenboschBarony struct {
	feud.BaseBarony
}

var BShertogenbosch斯海尔托亨博思 feud.Barony = &斯海尔托亨博思ShertogenboschBarony{}

func init() {
	f := BShertogenbosch斯海尔托亨博思.(*斯海尔托亨博思ShertogenboschBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shertogenbosch",
		TitleName: "斯海尔托亨博思",
		TitleCode: "b_shertogenbosch",
	}
}
