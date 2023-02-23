package york

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟斯克ThirskBarony struct {
	feud.BaseBarony
}

var BThirsk瑟斯克 feud.Barony = &瑟斯克ThirskBarony{}

func init() {
	f := BThirsk瑟斯克.(*瑟斯克ThirskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thirsk",
		TitleName: "瑟斯克",
		TitleCode: "b_thirsk",
	}
}
