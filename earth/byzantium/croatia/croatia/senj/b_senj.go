package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尼SenjBarony struct {
	feud.BaseBarony
}

var BSenj塞尼 feud.Barony = &塞尼SenjBarony{}

func init() {
    f := BSenj塞尼.(*塞尼SenjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "senj",
		TitleName: "塞尼",
		TitleCode: "b_senj",
	}
}
