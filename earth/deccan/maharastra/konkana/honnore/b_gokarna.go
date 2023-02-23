package honnore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿羯罗拏GokarnaBarony struct {
	feud.BaseBarony
}

var BGokarna瞿羯罗拏 feud.Barony = &瞿羯罗拏GokarnaBarony{}

func init() {
	f := BGokarna瞿羯罗拏.(*瞿羯罗拏GokarnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gokarna",
		TitleName: "瞿羯罗拏",
		TitleCode: "b_gokarna",
	}
}
