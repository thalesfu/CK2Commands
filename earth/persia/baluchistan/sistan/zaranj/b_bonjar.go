package zaranj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尼亚尔BonjarBarony struct {
	feud.BaseBarony
}

var BBonjar博尼亚尔 feud.Barony = &博尼亚尔BonjarBarony{}

func init() {
    f := BBonjar博尼亚尔.(*博尼亚尔BonjarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bonjar",
		TitleName: "博尼亚尔",
		TitleCode: "b_bonjar",
	}
}
