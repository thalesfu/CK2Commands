package izborsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图佩TupyBarony struct {
	feud.BaseBarony
}

var BTupy图佩 feud.Barony = &图佩TupyBarony{}

func init() {
    f := BTupy图佩.(*图佩TupyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tupy",
		TitleName: "图佩",
		TitleCode: "b_tupy",
	}
}
