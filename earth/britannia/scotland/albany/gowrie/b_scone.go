package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯昆SconeBarony struct {
	feud.BaseBarony
}

var BScone斯昆 feud.Barony = &斯昆SconeBarony{}

func init() {
    f := BScone斯昆.(*斯昆SconeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scone",
		TitleName: "斯昆",
		TitleCode: "b_scone",
	}
}
