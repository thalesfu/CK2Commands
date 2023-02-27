package dunhuang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玉门泉YumenquanBarony struct {
	feud.BaseBarony
}

var BYumenquan玉门泉 feud.Barony = &玉门泉YumenquanBarony{}

func init() {
    f := BYumenquan玉门泉.(*玉门泉YumenquanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yumenquan",
		TitleName: "玉门泉",
		TitleCode: "b_yumenquan",
	}
}
