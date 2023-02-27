package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉戈依RagaiBarony struct {
	feud.BaseBarony
}

var BRagai拉戈依 feud.Barony = &拉戈依RagaiBarony{}

func init() {
    f := BRagai拉戈依.(*拉戈依RagaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ragai",
		TitleName: "拉戈依",
		TitleCode: "b_ragai",
	}
}
