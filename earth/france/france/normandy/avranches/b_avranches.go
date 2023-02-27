package avranches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫朗什AvranchesBarony struct {
	feud.BaseBarony
}

var BAvranches阿夫朗什 feud.Barony = &阿夫朗什AvranchesBarony{}

func init() {
    f := BAvranches阿夫朗什.(*阿夫朗什AvranchesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avranches",
		TitleName: "阿夫朗什",
		TitleCode: "b_avranches",
	}
}
