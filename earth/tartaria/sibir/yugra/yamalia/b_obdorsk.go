package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鄂毕多尔斯克ObdorskBarony struct {
	feud.BaseBarony
}

var BObdorsk鄂毕多尔斯克 feud.Barony = &鄂毕多尔斯克ObdorskBarony{}

func init() {
	f := BObdorsk鄂毕多尔斯克.(*鄂毕多尔斯克ObdorskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "obdorsk",
		TitleName: "鄂毕多尔斯克",
		TitleCode: "b_obdorsk",
	}
}
