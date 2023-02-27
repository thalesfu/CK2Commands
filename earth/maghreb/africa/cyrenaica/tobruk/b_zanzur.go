package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赞祖尔ZanzurBarony struct {
	feud.BaseBarony
}

var BZanzur赞祖尔 feud.Barony = &赞祖尔ZanzurBarony{}

func init() {
    f := BZanzur赞祖尔.(*赞祖尔ZanzurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zanzur",
		TitleName: "赞祖尔",
		TitleCode: "b_zanzur",
	}
}
