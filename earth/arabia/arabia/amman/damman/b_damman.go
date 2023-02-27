package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达曼DammanBarony struct {
	feud.BaseBarony
}

var BDamman达曼 feud.Barony = &达曼DammanBarony{}

func init() {
    f := BDamman达曼.(*达曼DammanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damman",
		TitleName: "达曼",
		TitleCode: "b_damman",
	}
}
