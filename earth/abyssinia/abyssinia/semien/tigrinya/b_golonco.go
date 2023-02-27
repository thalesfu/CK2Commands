package tigrinya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哥隆可GoloncoBarony struct {
	feud.BaseBarony
}

var BGolonco哥隆可 feud.Barony = &哥隆可GoloncoBarony{}

func init() {
    f := BGolonco哥隆可.(*哥隆可GoloncoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "golonco",
		TitleName: "哥隆可",
		TitleCode: "b_golonco",
	}
}
