package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 初阿坦ChuartanBarony struct {
	feud.BaseBarony
}

var BChuartan初阿坦 feud.Barony = &初阿坦ChuartanBarony{}

func init() {
    f := BChuartan初阿坦.(*初阿坦ChuartanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chuartan",
		TitleName: "初阿坦",
		TitleCode: "b_chuartan",
	}
}
