package somnath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提伐陀DelvadaBarony struct {
	feud.BaseBarony
}

var BDelvada提伐陀 feud.Barony = &提伐陀DelvadaBarony{}

func init() {
	f := BDelvada提伐陀.(*提伐陀DelvadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "delvada",
		TitleName: "提伐陀",
		TitleCode: "b_delvada",
	}
}
