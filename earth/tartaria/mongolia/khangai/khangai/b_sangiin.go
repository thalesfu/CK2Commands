package khangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑根SangiinBarony struct {
	feud.BaseBarony
}

var BSangiin桑根 feud.Barony = &桑根SangiinBarony{}

func init() {
    f := BSangiin桑根.(*桑根SangiinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangiin",
		TitleName: "桑根",
		TitleCode: "b_sangiin",
	}
}
