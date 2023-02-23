package cherchen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 末城MochengBarony struct {
	feud.BaseBarony
}

var BMocheng末城 feud.Barony = &末城MochengBarony{}

func init() {
	f := BMocheng末城.(*末城MochengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mocheng",
		TitleName: "末城",
		TitleCode: "b_mocheng",
	}
}
