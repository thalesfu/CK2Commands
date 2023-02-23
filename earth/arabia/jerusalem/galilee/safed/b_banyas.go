package safed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尼亚斯BanyasBarony struct {
	feud.BaseBarony
}

var BBanyas巴尼亚斯 feud.Barony = &巴尼亚斯BanyasBarony{}

func init() {
	f := BBanyas巴尼亚斯.(*巴尼亚斯BanyasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banyas",
		TitleName: "巴尼亚斯",
		TitleCode: "b_banyas",
	}
}
