package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼亚波尔NeapolBarony struct {
	feud.BaseBarony
}

var BNeapol尼亚波尔 feud.Barony = &尼亚波尔NeapolBarony{}

func init() {
    f := BNeapol尼亚波尔.(*尼亚波尔NeapolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neapol",
		TitleName: "尼亚波尔",
		TitleCode: "b_neapol",
	}
}
