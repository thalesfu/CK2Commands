package turkestan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨克绍利斯基SaksaulskiyBarony struct {
	feud.BaseBarony
}

var BSaksaulskiy萨克绍利斯基 feud.Barony = &萨克绍利斯基SaksaulskiyBarony{}

func init() {
    f := BSaksaulskiy萨克绍利斯基.(*萨克绍利斯基SaksaulskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saksaulskiy",
		TitleName: "萨克绍利斯基",
		TitleCode: "b_saksaulskiy",
	}
}
