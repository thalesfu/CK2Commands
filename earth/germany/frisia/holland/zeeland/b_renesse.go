package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷讷瑟RenesseBarony struct {
	feud.BaseBarony
}

var BRenesse雷讷瑟 feud.Barony = &雷讷瑟RenesseBarony{}

func init() {
	f := BRenesse雷讷瑟.(*雷讷瑟RenesseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "renesse",
		TitleName: "雷讷瑟",
		TitleCode: "b_renesse",
	}
}
