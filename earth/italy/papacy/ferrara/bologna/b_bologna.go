package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博洛尼亚BolognaBarony struct {
	feud.BaseBarony
}

var BBologna博洛尼亚 feud.Barony = &博洛尼亚BolognaBarony{}

func init() {
	f := BBologna博洛尼亚.(*博洛尼亚BolognaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bologna",
		TitleName: "博洛尼亚",
		TitleCode: "b_bologna",
	}
}
