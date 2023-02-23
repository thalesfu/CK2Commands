package brycheiniog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔加斯TalgarthBarony struct {
	feud.BaseBarony
}

var BTalgarth塔尔加斯 feud.Barony = &塔尔加斯TalgarthBarony{}

func init() {
	f := BTalgarth塔尔加斯.(*塔尔加斯TalgarthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talgarth",
		TitleName: "塔尔加斯",
		TitleCode: "b_talgarth",
	}
}
