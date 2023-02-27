package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚沙尔塔YashaltaBarony struct {
	feud.BaseBarony
}

var BYashalta亚沙尔塔 feud.Barony = &亚沙尔塔YashaltaBarony{}

func init() {
    f := BYashalta亚沙尔塔.(*亚沙尔塔YashaltaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yashalta",
		TitleName: "亚沙尔塔",
		TitleCode: "b_yashalta",
	}
}
