package krems

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃根堡EggenburgBarony struct {
	feud.BaseBarony
}

var BEggenburg埃根堡 feud.Barony = &埃根堡EggenburgBarony{}

func init() {
    f := BEggenburg埃根堡.(*埃根堡EggenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eggenburg",
		TitleName: "埃根堡",
		TitleCode: "b_eggenburg",
	}
}
