package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑迪SandayBarony struct {
	feud.BaseBarony
}

var BSanday桑迪 feud.Barony = &桑迪SandayBarony{}

func init() {
	f := BSanday桑迪.(*桑迪SandayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanday",
		TitleName: "桑迪",
		TitleCode: "b_sanday",
	}
}
