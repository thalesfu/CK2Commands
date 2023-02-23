package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滕斯贝格TonsbergBarony struct {
	feud.BaseBarony
}

var BTonsberg滕斯贝格 feud.Barony = &滕斯贝格TonsbergBarony{}

func init() {
	f := BTonsberg滕斯贝格.(*滕斯贝格TonsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tonsberg",
		TitleName: "滕斯贝格",
		TitleCode: "b_tonsberg",
	}
}
