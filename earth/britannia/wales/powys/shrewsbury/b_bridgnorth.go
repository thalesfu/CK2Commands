package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里奇诺斯BridgnorthBarony struct {
	feud.BaseBarony
}

var BBridgnorth布里奇诺斯 feud.Barony = &布里奇诺斯BridgnorthBarony{}

func init() {
	f := BBridgnorth布里奇诺斯.(*布里奇诺斯BridgnorthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bridgnorth",
		TitleName: "布里奇诺斯",
		TitleCode: "b_bridgnorth",
	}
}
