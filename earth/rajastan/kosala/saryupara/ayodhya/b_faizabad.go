package ayodhya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法扎巴德FaizabadBarony struct {
	feud.BaseBarony
}

var BFaizabad法扎巴德 feud.Barony = &法扎巴德FaizabadBarony{}

func init() {
    f := BFaizabad法扎巴德.(*法扎巴德FaizabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faizabad",
		TitleName: "法扎巴德",
		TitleCode: "b_faizabad",
	}
}
