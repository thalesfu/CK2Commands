package quzdar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嫡鸠牟利TekkumuriBarony struct {
	feud.BaseBarony
}

var BTekkumuri嫡鸠牟利 feud.Barony = &嫡鸠牟利TekkumuriBarony{}

func init() {
	f := BTekkumuri嫡鸠牟利.(*嫡鸠牟利TekkumuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tekkumuri",
		TitleName: "嫡鸠牟利",
		TitleCode: "b_tekkumuri",
	}
}
