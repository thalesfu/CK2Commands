package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗莫朗坦RomorantinBarony struct {
	feud.BaseBarony
}

var BRomorantin罗莫朗坦 feud.Barony = &罗莫朗坦RomorantinBarony{}

func init() {
    f := BRomorantin罗莫朗坦.(*罗莫朗坦RomorantinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "romorantin",
		TitleName: "罗莫朗坦",
		TitleCode: "b_romorantin",
	}
}
