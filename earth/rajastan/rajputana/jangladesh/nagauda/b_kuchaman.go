package nagauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古贾曼KuchamanBarony struct {
	feud.BaseBarony
}

var BKuchaman古贾曼 feud.Barony = &古贾曼KuchamanBarony{}

func init() {
	f := BKuchaman古贾曼.(*古贾曼KuchamanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuchaman",
		TitleName: "古贾曼",
		TitleCode: "b_kuchaman",
	}
}
