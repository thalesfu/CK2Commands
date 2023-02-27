package tuul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嘎丘尔特GachuurtBarony struct {
	feud.BaseBarony
}

var BGachuurt嘎丘尔特 feud.Barony = &嘎丘尔特GachuurtBarony{}

func init() {
    f := BGachuurt嘎丘尔特.(*嘎丘尔特GachuurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gachuurt",
		TitleName: "嘎丘尔特",
		TitleCode: "b_gachuurt",
	}
}
