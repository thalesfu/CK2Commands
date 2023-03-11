package desht-i-kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利曼LymanBarony struct {
	feud.BaseBarony
}

var BLyman利曼 feud.Barony = &利曼LymanBarony{}

func init() {
    f := BLyman利曼.(*利曼LymanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyman",
		TitleName: "利曼",
		TitleCode: "b_lyman",
	}
}
