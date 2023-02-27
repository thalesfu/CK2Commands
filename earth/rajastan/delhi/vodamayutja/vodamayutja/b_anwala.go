package vodamayutja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦罗AnwalaBarony struct {
	feud.BaseBarony
}

var BAnwala阿瓦罗 feud.Barony = &阿瓦罗AnwalaBarony{}

func init() {
    f := BAnwala阿瓦罗.(*阿瓦罗AnwalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anwala",
		TitleName: "阿瓦罗",
		TitleCode: "b_anwala",
	}
}
