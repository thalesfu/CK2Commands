package bhakkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珀格尔BhakkarBarony struct {
	feud.BaseBarony
}

var BBhakkar珀格尔 feud.Barony = &珀格尔BhakkarBarony{}

func init() {
    f := BBhakkar珀格尔.(*珀格尔BhakkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhakkar",
		TitleName: "珀格尔",
		TitleCode: "b_bhakkar",
	}
}
