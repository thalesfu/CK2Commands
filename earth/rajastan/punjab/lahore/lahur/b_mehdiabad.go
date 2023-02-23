package lahur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈赫迪耶巴德MehdiabadBarony struct {
	feud.BaseBarony
}

var BMehdiabad迈赫迪耶巴德 feud.Barony = &迈赫迪耶巴德MehdiabadBarony{}

func init() {
	f := BMehdiabad迈赫迪耶巴德.(*迈赫迪耶巴德MehdiabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mehdiabad",
		TitleName: "迈赫迪耶巴德",
		TitleCode: "b_mehdiabad",
	}
}
