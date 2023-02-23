package damoh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼牟罗NimoraBarony struct {
	feud.BaseBarony
}

var BNimora尼牟罗 feud.Barony = &尼牟罗NimoraBarony{}

func init() {
	f := BNimora尼牟罗.(*尼牟罗NimoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nimora",
		TitleName: "尼牟罗",
		TitleCode: "b_nimora",
	}
}
