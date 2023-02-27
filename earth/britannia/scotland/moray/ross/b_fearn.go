package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗恩FearnBarony struct {
	feud.BaseBarony
}

var BFearn弗恩 feud.Barony = &弗恩FearnBarony{}

func init() {
    f := BFearn弗恩.(*弗恩FearnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fearn",
		TitleName: "弗恩",
		TitleCode: "b_fearn",
	}
}
