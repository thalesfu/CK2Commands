package shish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦西斯VasissBarony struct {
	feud.BaseBarony
}

var BVasiss瓦西斯 feud.Barony = &瓦西斯VasissBarony{}

func init() {
    f := BVasiss瓦西斯.(*瓦西斯VasissBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vasiss",
		TitleName: "瓦西斯",
		TitleCode: "b_vasiss",
	}
}
