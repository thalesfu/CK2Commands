package kotthasara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿羯罗拏GokannaBarony struct {
	feud.BaseBarony
}

var BGokanna瞿羯罗拏 feud.Barony = &瞿羯罗拏GokannaBarony{}

func init() {
    f := BGokanna瞿羯罗拏.(*瞿羯罗拏GokannaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gokanna",
		TitleName: "瞿羯罗拏",
		TitleCode: "b_gokanna",
	}
}
