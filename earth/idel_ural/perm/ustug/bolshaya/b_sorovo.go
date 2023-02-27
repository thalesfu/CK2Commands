package bolshaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索罗沃SorovoBarony struct {
	feud.BaseBarony
}

var BSorovo索罗沃 feud.Barony = &索罗沃SorovoBarony{}

func init() {
    f := BSorovo索罗沃.(*索罗沃SorovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorovo",
		TitleName: "索罗沃",
		TitleCode: "b_sorovo",
	}
}
