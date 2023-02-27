package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈默斯胡斯HammershusBarony struct {
	feud.BaseBarony
}

var BHammershus哈默斯胡斯 feud.Barony = &哈默斯胡斯HammershusBarony{}

func init() {
    f := BHammershus哈默斯胡斯.(*哈默斯胡斯HammershusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hammershus",
		TitleName: "哈默斯胡斯",
		TitleCode: "b_hammershus",
	}
}
