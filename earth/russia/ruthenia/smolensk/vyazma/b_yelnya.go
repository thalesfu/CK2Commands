package vyazma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶利尼亚YelnyaBarony struct {
	feud.BaseBarony
}

var BYelnya叶利尼亚 feud.Barony = &叶利尼亚YelnyaBarony{}

func init() {
	f := BYelnya叶利尼亚.(*叶利尼亚YelnyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yelnya",
		TitleName: "叶利尼亚",
		TitleCode: "b_yelnya",
	}
}
