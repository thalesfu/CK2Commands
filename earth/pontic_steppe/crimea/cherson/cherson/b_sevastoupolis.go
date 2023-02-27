package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞瓦斯图波利斯SevastoupolisBarony struct {
	feud.BaseBarony
}

var BSevastoupolis塞瓦斯图波利斯 feud.Barony = &塞瓦斯图波利斯SevastoupolisBarony{}

func init() {
    f := BSevastoupolis塞瓦斯图波利斯.(*塞瓦斯图波利斯SevastoupolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sevastoupolis",
		TitleName: "塞瓦斯图波利斯",
		TitleCode: "b_sevastoupolis",
	}
}
