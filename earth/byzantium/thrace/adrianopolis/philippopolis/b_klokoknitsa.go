package philippopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗克尼萨KlokoknitsaBarony struct {
	feud.BaseBarony
}

var BKlokoknitsa克罗克尼萨 feud.Barony = &克罗克尼萨KlokoknitsaBarony{}

func init() {
    f := BKlokoknitsa克罗克尼萨.(*克罗克尼萨KlokoknitsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klokoknitsa",
		TitleName: "克罗克尼萨",
		TitleCode: "b_klokoknitsa",
	}
}
