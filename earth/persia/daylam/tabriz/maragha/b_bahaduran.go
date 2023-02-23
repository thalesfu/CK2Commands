package maragha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴哈杜兰BahaduranBarony struct {
	feud.BaseBarony
}

var BBahaduran巴哈杜兰 feud.Barony = &巴哈杜兰BahaduranBarony{}

func init() {
	f := BBahaduran巴哈杜兰.(*巴哈杜兰BahaduranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahaduran",
		TitleName: "巴哈杜兰",
		TitleCode: "b_bahaduran",
	}
}
