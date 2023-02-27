package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯艾因ResainaBarony struct {
	feud.BaseBarony
}

var BResaina拉斯艾因 feud.Barony = &拉斯艾因ResainaBarony{}

func init() {
    f := BResaina拉斯艾因.(*拉斯艾因ResainaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "resaina",
		TitleName: "拉斯艾因",
		TitleCode: "b_resaina",
	}
}
