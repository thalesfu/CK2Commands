package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥古斯朗AlgutsrumBarony struct {
	feud.BaseBarony
}

var BAlgutsrum奥古斯朗 feud.Barony = &奥古斯朗AlgutsrumBarony{}

func init() {
	f := BAlgutsrum奥古斯朗.(*奥古斯朗AlgutsrumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "algutsrum",
		TitleName: "奥古斯朗",
		TitleCode: "b_algutsrum",
	}
}
