package asyut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜特哈利拉夫BeitkhallafBarony struct {
	feud.BaseBarony
}

var BBeitkhallaf拜特哈利拉夫 feud.Barony = &拜特哈利拉夫BeitkhallafBarony{}

func init() {
    f := BBeitkhallaf拜特哈利拉夫.(*拜特哈利拉夫BeitkhallafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beitkhallaf",
		TitleName: "拜特哈利拉夫",
		TitleCode: "b_beitkhallaf",
	}
}
