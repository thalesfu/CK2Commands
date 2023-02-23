package philippopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴奇科沃BachkovoBarony struct {
	feud.BaseBarony
}

var BBachkovo巴奇科沃 feud.Barony = &巴奇科沃BachkovoBarony{}

func init() {
	f := BBachkovo巴奇科沃.(*巴奇科沃BachkovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bachkovo",
		TitleName: "巴奇科沃",
		TitleCode: "b_bachkovo",
	}
}
