package ceredigion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯赖伦AberaeronBarony struct {
	feud.BaseBarony
}

var BAberaeron阿伯赖伦 feud.Barony = &阿伯赖伦AberaeronBarony{}

func init() {
    f := BAberaeron阿伯赖伦.(*阿伯赖伦AberaeronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aberaeron",
		TitleName: "阿伯赖伦",
		TitleCode: "b_aberaeron",
	}
}
