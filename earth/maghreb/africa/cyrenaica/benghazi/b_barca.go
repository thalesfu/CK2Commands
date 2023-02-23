package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔切BarcaBarony struct {
	feud.BaseBarony
}

var BBarca巴尔切 feud.Barony = &巴尔切BarcaBarony{}

func init() {
	f := BBarca巴尔切.(*巴尔切BarcaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barca",
		TitleName: "巴尔切",
		TitleCode: "b_barca",
	}
}
