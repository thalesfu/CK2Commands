package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察甘_阿曼TsaganamanBarony struct {
	feud.BaseBarony
}

var BTsaganaman察甘_阿曼 feud.Barony = &察甘_阿曼TsaganamanBarony{}

func init() {
    f := BTsaganaman察甘_阿曼.(*察甘_阿曼TsaganamanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsaganaman",
		TitleName: "察甘_阿曼",
		TitleCode: "b_tsaganaman",
	}
}
