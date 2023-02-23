package balkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏尔赫科塔尔SurkhkotalBarony struct {
	feud.BaseBarony
}

var BSurkhkotal苏尔赫科塔尔 feud.Barony = &苏尔赫科塔尔SurkhkotalBarony{}

func init() {
	f := BSurkhkotal苏尔赫科塔尔.(*苏尔赫科塔尔SurkhkotalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surkhkotal",
		TitleName: "苏尔赫科塔尔",
		TitleCode: "b_surkhkotal",
	}
}
