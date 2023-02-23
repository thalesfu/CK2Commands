package adrianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯科洛特SkalotheBarony struct {
	feud.BaseBarony
}

var BSkalothe斯科洛特 feud.Barony = &斯科洛特SkalotheBarony{}

func init() {
	f := BSkalothe斯科洛特.(*斯科洛特SkalotheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skalothe",
		TitleName: "斯科洛特",
		TitleCode: "b_skalothe",
	}
}
