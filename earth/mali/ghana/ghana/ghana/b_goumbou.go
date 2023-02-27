package ghana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡布GoumbouBarony struct {
	feud.BaseBarony
}

var BGoumbou贡布 feud.Barony = &贡布GoumbouBarony{}

func init() {
    f := BGoumbou贡布.(*贡布GoumbouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goumbou",
		TitleName: "贡布",
		TitleCode: "b_goumbou",
	}
}
