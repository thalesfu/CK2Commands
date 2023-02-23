package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔谢巴BeershebBarony struct {
	feud.BaseBarony
}

var BBeersheb贝尔谢巴 feud.Barony = &贝尔谢巴BeershebBarony{}

func init() {
	f := BBeersheb贝尔谢巴.(*贝尔谢巴BeershebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beersheb",
		TitleName: "贝尔谢巴",
		TitleCode: "b_beersheb",
	}
}
