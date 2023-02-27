package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴德泽本BardzerbenBarony struct {
	feud.BaseBarony
}

var BBardzerben巴德泽本 feud.Barony = &巴德泽本BardzerbenBarony{}

func init() {
    f := BBardzerben巴德泽本.(*巴德泽本BardzerbenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bardzerben",
		TitleName: "巴德泽本",
		TitleCode: "b_bardzerben",
	}
}
