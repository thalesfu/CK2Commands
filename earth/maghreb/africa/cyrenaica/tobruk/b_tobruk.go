package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图卜鲁格TobrukBarony struct {
	feud.BaseBarony
}

var BTobruk图卜鲁格 feud.Barony = &图卜鲁格TobrukBarony{}

func init() {
    f := BTobruk图卜鲁格.(*图卜鲁格TobrukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tobruk",
		TitleName: "图卜鲁格",
		TitleCode: "b_tobruk",
	}
}
