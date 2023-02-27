package amastris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马斯特里斯AmastrisBarony struct {
	feud.BaseBarony
}

var BAmastris阿马斯特里斯 feud.Barony = &阿马斯特里斯AmastrisBarony{}

func init() {
    f := BAmastris阿马斯特里斯.(*阿马斯特里斯AmastrisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amastris",
		TitleName: "阿马斯特里斯",
		TitleCode: "b_amastris",
	}
}
