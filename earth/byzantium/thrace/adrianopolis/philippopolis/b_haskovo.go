package philippopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈斯科沃HaskovoBarony struct {
	feud.BaseBarony
}

var BHaskovo哈斯科沃 feud.Barony = &哈斯科沃HaskovoBarony{}

func init() {
	f := BHaskovo哈斯科沃.(*哈斯科沃HaskovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haskovo",
		TitleName: "哈斯科沃",
		TitleCode: "b_haskovo",
	}
}
