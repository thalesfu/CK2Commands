package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰纳兰ChenaranBarony struct {
	feud.BaseBarony
}

var BChenaran恰纳兰 feud.Barony = &恰纳兰ChenaranBarony{}

func init() {
	f := BChenaran恰纳兰.(*恰纳兰ChenaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chenaran",
		TitleName: "恰纳兰",
		TitleCode: "b_chenaran",
	}
}
