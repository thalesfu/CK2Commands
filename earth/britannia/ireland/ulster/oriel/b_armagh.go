package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马ArmaghBarony struct {
	feud.BaseBarony
}

var BArmagh阿马 feud.Barony = &阿马ArmaghBarony{}

func init() {
    f := BArmagh阿马.(*阿马ArmaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "armagh",
		TitleName: "阿马",
		TitleCode: "b_armagh",
	}
}
