package pongau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔姆斯韦格TaemswichBarony struct {
	feud.BaseBarony
}

var BTaemswich塔姆斯韦格 feud.Barony = &塔姆斯韦格TaemswichBarony{}

func init() {
	f := BTaemswich塔姆斯韦格.(*塔姆斯韦格TaemswichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taemswich",
		TitleName: "塔姆斯韦格",
		TitleCode: "b_taemswich",
	}
}
