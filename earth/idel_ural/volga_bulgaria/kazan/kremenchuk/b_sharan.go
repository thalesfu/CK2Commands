package kremenchuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙兰SharanBarony struct {
	feud.BaseBarony
}

var BSharan沙兰 feud.Barony = &沙兰SharanBarony{}

func init() {
    f := BSharan沙兰.(*沙兰SharanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharan",
		TitleName: "沙兰",
		TitleCode: "b_sharan",
	}
}
