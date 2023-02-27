package irtek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯里利BoriliBarony struct {
	feud.BaseBarony
}

var BBorili伯里利 feud.Barony = &伯里利BoriliBarony{}

func init() {
    f := BBorili伯里利.(*伯里利BoriliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borili",
		TitleName: "伯里利",
		TitleCode: "b_borili",
	}
}
