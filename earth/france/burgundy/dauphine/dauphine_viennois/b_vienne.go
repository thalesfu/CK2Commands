package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维埃纳VienneBarony struct {
	feud.BaseBarony
}

var BVienne维埃纳 feud.Barony = &维埃纳VienneBarony{}

func init() {
    f := BVienne维埃纳.(*维埃纳VienneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vienne",
		TitleName: "维埃纳",
		TitleCode: "b_vienne",
	}
}
