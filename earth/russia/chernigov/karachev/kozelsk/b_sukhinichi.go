package kozelsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏希尼奇SukhinichiBarony struct {
	feud.BaseBarony
}

var BSukhinichi苏希尼奇 feud.Barony = &苏希尼奇SukhinichiBarony{}

func init() {
	f := BSukhinichi苏希尼奇.(*苏希尼奇SukhinichiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sukhinichi",
		TitleName: "苏希尼奇",
		TitleCode: "b_sukhinichi",
	}
}
