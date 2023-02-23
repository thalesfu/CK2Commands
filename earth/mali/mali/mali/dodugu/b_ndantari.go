package dodugu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩丹塔里NdantariBarony struct {
	feud.BaseBarony
}

var BNdantari恩丹塔里 feud.Barony = &恩丹塔里NdantariBarony{}

func init() {
	f := BNdantari恩丹塔里.(*恩丹塔里NdantariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ndantari",
		TitleName: "恩丹塔里",
		TitleCode: "b_ndantari",
	}
}
