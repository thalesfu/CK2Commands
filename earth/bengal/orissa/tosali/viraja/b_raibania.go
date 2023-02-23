package viraja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉巴尼亚RaibaniaBarony struct {
	feud.BaseBarony
}

var BRaibania拉巴尼亚 feud.Barony = &拉巴尼亚RaibaniaBarony{}

func init() {
	f := BRaibania拉巴尼亚.(*拉巴尼亚RaibaniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raibania",
		TitleName: "拉巴尼亚",
		TitleCode: "b_raibania",
	}
}
