package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉德AradenBarony struct {
	feud.BaseBarony
}

var BAraden阿拉德 feud.Barony = &阿拉德AradenBarony{}

func init() {
	f := BAraden阿拉德.(*阿拉德AradenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "araden",
		TitleName: "阿拉德",
		TitleCode: "b_araden",
	}
}
