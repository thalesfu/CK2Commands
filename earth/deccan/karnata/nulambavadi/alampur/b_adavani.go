package alampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿达瓦尼AdavaniBarony struct {
	feud.BaseBarony
}

var BAdavani阿达瓦尼 feud.Barony = &阿达瓦尼AdavaniBarony{}

func init() {
	f := BAdavani阿达瓦尼.(*阿达瓦尼AdavaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adavani",
		TitleName: "阿达瓦尼",
		TitleCode: "b_adavani",
	}
}
