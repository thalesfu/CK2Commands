package anhalt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维滕贝格Wittenberg_zeistBarony struct {
	feud.BaseBarony
}

var BWittenberg_zeist维滕贝格 feud.Barony = &维滕贝格Wittenberg_zeistBarony{}

func init() {
    f := BWittenberg_zeist维滕贝格.(*维滕贝格Wittenberg_zeistBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wittenberg_zeist",
		TitleName: "维滕贝格",
		TitleCode: "b_wittenberg_zeist",
	}
}
