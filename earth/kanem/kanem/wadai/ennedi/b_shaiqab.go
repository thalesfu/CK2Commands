package ennedi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢加卜ShaiqabBarony struct {
	feud.BaseBarony
}

var BShaiqab谢加卜 feud.Barony = &谢加卜ShaiqabBarony{}

func init() {
	f := BShaiqab谢加卜.(*谢加卜ShaiqabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shaiqab",
		TitleName: "谢加卜",
		TitleCode: "b_shaiqab",
	}
}
