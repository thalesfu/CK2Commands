package sortavala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁斯克阿拉RuskealaBarony struct {
	feud.BaseBarony
}

var BRuskeala鲁斯克阿拉 feud.Barony = &鲁斯克阿拉RuskealaBarony{}

func init() {
    f := BRuskeala鲁斯克阿拉.(*鲁斯克阿拉RuskealaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruskeala",
		TitleName: "鲁斯克阿拉",
		TitleCode: "b_ruskeala",
	}
}
