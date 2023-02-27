package laukaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉韦西SaraavesiBarony struct {
	feud.BaseBarony
}

var BSaraavesi萨拉韦西 feud.Barony = &萨拉韦西SaraavesiBarony{}

func init() {
    f := BSaraavesi萨拉韦西.(*萨拉韦西SaraavesiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saraavesi",
		TitleName: "萨拉韦西",
		TitleCode: "b_saraavesi",
	}
}
