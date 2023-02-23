package ephesos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉莱斯TrallesBarony struct {
	feud.BaseBarony
}

var BTralles特拉莱斯 feud.Barony = &特拉莱斯TrallesBarony{}

func init() {
	f := BTralles特拉莱斯.(*特拉莱斯TrallesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tralles",
		TitleName: "特拉莱斯",
		TitleCode: "b_tralles",
	}
}
