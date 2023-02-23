package corfu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德瑞帕农DrepaneBarony struct {
	feud.BaseBarony
}

var BDrepane德瑞帕农 feud.Barony = &德瑞帕农DrepaneBarony{}

func init() {
	f := BDrepane德瑞帕农.(*德瑞帕农DrepaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drepane",
		TitleName: "德瑞帕农",
		TitleCode: "b_drepane",
	}
}
