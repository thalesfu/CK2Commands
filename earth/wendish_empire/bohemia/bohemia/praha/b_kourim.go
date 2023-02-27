package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考日姆KourimBarony struct {
	feud.BaseBarony
}

var BKourim考日姆 feud.Barony = &考日姆KourimBarony{}

func init() {
    f := BKourim考日姆.(*考日姆KourimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kourim",
		TitleName: "考日姆",
		TitleCode: "b_kourim",
	}
}
