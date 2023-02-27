package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰斯ReimsBarony struct {
	feud.BaseBarony
}

var BReims兰斯 feud.Barony = &兰斯ReimsBarony{}

func init() {
    f := BReims兰斯.(*兰斯ReimsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reims",
		TitleName: "兰斯",
		TitleCode: "b_reims",
	}
}
