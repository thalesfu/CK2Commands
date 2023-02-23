package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 撒摩拿ZaimonaBarony struct {
	feud.BaseBarony
}

var BZaimona撒摩拿 feud.Barony = &撒摩拿ZaimonaBarony{}

func init() {
	f := BZaimona撒摩拿.(*撒摩拿ZaimonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaimona",
		TitleName: "撒摩拿",
		TitleCode: "b_zaimona",
	}
}
