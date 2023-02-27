package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马劳毛罗什西盖特MaramarosszigetBarony struct {
	feud.BaseBarony
}

var BMaramarossziget马劳毛罗什西盖特 feud.Barony = &马劳毛罗什西盖特MaramarosszigetBarony{}

func init() {
    f := BMaramarossziget马劳毛罗什西盖特.(*马劳毛罗什西盖特MaramarosszigetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maramarossziget",
		TitleName: "马劳毛罗什西盖特",
		TitleCode: "b_maramarossziget",
	}
}
