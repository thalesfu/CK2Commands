package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒_阿斯克尔KyzylaskerBarony struct {
	feud.BaseBarony
}

var BKyzylasker克孜勒_阿斯克尔 feud.Barony = &克孜勒_阿斯克尔KyzylaskerBarony{}

func init() {
    f := BKyzylasker克孜勒_阿斯克尔.(*克孜勒_阿斯克尔KyzylaskerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzylasker",
		TitleName: "克孜勒_阿斯克尔",
		TitleCode: "b_kyzylasker",
	}
}
