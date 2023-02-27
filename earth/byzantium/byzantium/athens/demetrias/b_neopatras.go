package demetrias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼奥帕特拉斯NeopatrasBarony struct {
	feud.BaseBarony
}

var BNeopatras尼奥帕特拉斯 feud.Barony = &尼奥帕特拉斯NeopatrasBarony{}

func init() {
    f := BNeopatras尼奥帕特拉斯.(*尼奥帕特拉斯NeopatrasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neopatras",
		TitleName: "尼奥帕特拉斯",
		TitleCode: "b_neopatras",
	}
}
