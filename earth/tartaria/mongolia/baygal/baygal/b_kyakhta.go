package baygal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰克图KyakhtaBarony struct {
	feud.BaseBarony
}

var BKyakhta恰克图 feud.Barony = &恰克图KyakhtaBarony{}

func init() {
    f := BKyakhta恰克图.(*恰克图KyakhtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyakhta",
		TitleName: "恰克图",
		TitleCode: "b_kyakhta",
	}
}
