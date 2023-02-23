package zaysan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖拜尔ZubairBarony struct {
	feud.BaseBarony
}

var BZubair祖拜尔 feud.Barony = &祖拜尔ZubairBarony{}

func init() {
	f := BZubair祖拜尔.(*祖拜尔ZubairBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zubair",
		TitleName: "祖拜尔",
		TitleCode: "b_zubair",
	}
}
