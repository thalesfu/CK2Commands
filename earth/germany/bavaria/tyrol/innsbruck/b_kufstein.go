package innsbruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库夫施泰因KufsteinBarony struct {
	feud.BaseBarony
}

var BKufstein库夫施泰因 feud.Barony = &库夫施泰因KufsteinBarony{}

func init() {
    f := BKufstein库夫施泰因.(*库夫施泰因KufsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kufstein",
		TitleName: "库夫施泰因",
		TitleCode: "b_kufstein",
	}
}
