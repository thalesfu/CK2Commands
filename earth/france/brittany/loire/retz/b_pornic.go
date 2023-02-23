package retz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔尼克PornicBarony struct {
	feud.BaseBarony
}

var BPornic波尔尼克 feud.Barony = &波尔尼克PornicBarony{}

func init() {
	f := BPornic波尔尼克.(*波尔尼克PornicBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pornic",
		TitleName: "波尔尼克",
		TitleCode: "b_pornic",
	}
}
