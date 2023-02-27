package chortitza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃兹涅先斯克VosnesjenskBarony struct {
	feud.BaseBarony
}

var BVosnesjensk沃兹涅先斯克 feud.Barony = &沃兹涅先斯克VosnesjenskBarony{}

func init() {
    f := BVosnesjensk沃兹涅先斯克.(*沃兹涅先斯克VosnesjenskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vosnesjensk",
		TitleName: "沃兹涅先斯克",
		TitleCode: "b_vosnesjensk",
	}
}
