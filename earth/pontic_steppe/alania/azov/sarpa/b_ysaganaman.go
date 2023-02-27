package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察甘阿曼YsaganamanBarony struct {
	feud.BaseBarony
}

var BYsaganaman察甘阿曼 feud.Barony = &察甘阿曼YsaganamanBarony{}

func init() {
    f := BYsaganaman察甘阿曼.(*察甘阿曼YsaganamanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ysaganaman",
		TitleName: "察甘阿曼",
		TitleCode: "b_ysaganaman",
	}
}
