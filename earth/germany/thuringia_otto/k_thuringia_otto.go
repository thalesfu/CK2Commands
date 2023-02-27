package thuringia_otto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Thuringia_ottoKingdom interface {
    feud.Kingdom
}

type 图林根Thuringia_ottoKingdom struct {
	feud.BaseKingdom
}

var KThuringia_otto图林根 Thuringia_ottoKingdom = &图林根Thuringia_ottoKingdom{}

func init() {
	f := KThuringia_otto图林根.(*图林根Thuringia_ottoKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "thuringia_otto",
		TitleName: "图林根",
		TitleCode: "k_thuringia_otto",
		Dukes:  map[string]feud.Duke{},
	}

}
