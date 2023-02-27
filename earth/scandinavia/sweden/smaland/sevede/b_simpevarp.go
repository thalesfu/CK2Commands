package sevede

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛佩瓦普SimpevarpBarony struct {
	feud.BaseBarony
}

var BSimpevarp辛佩瓦普 feud.Barony = &辛佩瓦普SimpevarpBarony{}

func init() {
    f := BSimpevarp辛佩瓦普.(*辛佩瓦普SimpevarpBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "simpevarp",
		TitleName: "辛佩瓦普",
		TitleCode: "b_simpevarp",
	}
}
