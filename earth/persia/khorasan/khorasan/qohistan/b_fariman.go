package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法里曼FarimanBarony struct {
	feud.BaseBarony
}

var BFariman法里曼 feud.Barony = &法里曼FarimanBarony{}

func init() {
    f := BFariman法里曼.(*法里曼FarimanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fariman",
		TitleName: "法里曼",
		TitleCode: "b_fariman",
	}
}
