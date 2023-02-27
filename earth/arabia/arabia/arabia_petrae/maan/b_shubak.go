package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏巴克ShubakBarony struct {
	feud.BaseBarony
}

var BShubak苏巴克 feud.Barony = &苏巴克ShubakBarony{}

func init() {
    f := BShubak苏巴克.(*苏巴克ShubakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shubak",
		TitleName: "苏巴克",
		TitleCode: "b_shubak",
	}
}
