package lomzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特罗文卡OstrolekaBarony struct {
	feud.BaseBarony
}

var BOstroleka奥斯特罗文卡 feud.Barony = &奥斯特罗文卡OstrolekaBarony{}

func init() {
    f := BOstroleka奥斯特罗文卡.(*奥斯特罗文卡OstrolekaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostroleka",
		TitleName: "奥斯特罗文卡",
		TitleCode: "b_ostroleka",
	}
}
