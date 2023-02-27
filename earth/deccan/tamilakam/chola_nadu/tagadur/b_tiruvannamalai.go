package tagadur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提卢瓦那摩罗TiruvannamalaiBarony struct {
	feud.BaseBarony
}

var BTiruvannamalai提卢瓦那摩罗 feud.Barony = &提卢瓦那摩罗TiruvannamalaiBarony{}

func init() {
    f := BTiruvannamalai提卢瓦那摩罗.(*提卢瓦那摩罗TiruvannamalaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiruvannamalai",
		TitleName: "提卢瓦那摩罗",
		TitleCode: "b_tiruvannamalai",
	}
}
