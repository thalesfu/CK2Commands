package galatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫基索斯MikissosBarony struct {
	feud.BaseBarony
}

var BMikissos莫基索斯 feud.Barony = &莫基索斯MikissosBarony{}

func init() {
	f := BMikissos莫基索斯.(*莫基索斯MikissosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mikissos",
		TitleName: "莫基索斯",
		TitleCode: "b_mikissos",
	}
}
