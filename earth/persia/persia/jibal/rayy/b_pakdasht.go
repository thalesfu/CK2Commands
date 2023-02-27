package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕克达什特PakdashtBarony struct {
	feud.BaseBarony
}

var BPakdasht帕克达什特 feud.Barony = &帕克达什特PakdashtBarony{}

func init() {
    f := BPakdasht帕克达什特.(*帕克达什特PakdashtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pakdasht",
		TitleName: "帕克达什特",
		TitleCode: "b_pakdasht",
	}
}
