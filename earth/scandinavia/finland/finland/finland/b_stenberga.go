package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯腾贝里亚StenbergaBarony struct {
	feud.BaseBarony
}

var BStenberga斯腾贝里亚 feud.Barony = &斯腾贝里亚StenbergaBarony{}

func init() {
	f := BStenberga斯腾贝里亚.(*斯腾贝里亚StenbergaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stenberga",
		TitleName: "斯腾贝里亚",
		TitleCode: "b_stenberga",
	}
}
