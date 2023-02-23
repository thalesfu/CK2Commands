package bourbon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏维尼SouvignyBarony struct {
	feud.BaseBarony
}

var BSouvigny苏维尼 feud.Barony = &苏维尼SouvignyBarony{}

func init() {
	f := BSouvigny苏维尼.(*苏维尼SouvignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "souvigny",
		TitleName: "苏维尼",
		TitleCode: "b_souvigny",
	}
}
