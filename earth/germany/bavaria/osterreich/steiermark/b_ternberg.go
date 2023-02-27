package steiermark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滕贝格TernbergBarony struct {
	feud.BaseBarony
}

var BTernberg滕贝格 feud.Barony = &滕贝格TernbergBarony{}

func init() {
    f := BTernberg滕贝格.(*滕贝格TernbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ternberg",
		TitleName: "滕贝格",
		TitleCode: "b_ternberg",
	}
}
