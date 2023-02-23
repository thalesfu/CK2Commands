package geneve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托内ThonexBarony struct {
	feud.BaseBarony
}

var BThonex托内 feud.Barony = &托内ThonexBarony{}

func init() {
	f := BThonex托内.(*托内ThonexBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thonex",
		TitleName: "托内",
		TitleCode: "b_thonex",
	}
}
