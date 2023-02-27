package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下奥杰斯NizhodesBarony struct {
	feud.BaseBarony
}

var BNizhodes下奥杰斯 feud.Barony = &下奥杰斯NizhodesBarony{}

func init() {
    f := BNizhodes下奥杰斯.(*下奥杰斯NizhodesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizhodes",
		TitleName: "下奥杰斯",
		TitleCode: "b_nizhodes",
	}
}
