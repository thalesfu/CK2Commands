package pithapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达洛叉罗摩DraksharamaBarony struct {
	feud.BaseBarony
}

var BDraksharama达洛叉罗摩 feud.Barony = &达洛叉罗摩DraksharamaBarony{}

func init() {
	f := BDraksharama达洛叉罗摩.(*达洛叉罗摩DraksharamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "draksharama",
		TitleName: "达洛叉罗摩",
		TitleCode: "b_draksharama",
	}
}
