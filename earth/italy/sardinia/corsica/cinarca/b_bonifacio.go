package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尼法乔BonifacioBarony struct {
	feud.BaseBarony
}

var BBonifacio博尼法乔 feud.Barony = &博尼法乔BonifacioBarony{}

func init() {
    f := BBonifacio博尼法乔.(*博尼法乔BonifacioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bonifacio",
		TitleName: "博尼法乔",
		TitleCode: "b_bonifacio",
	}
}
