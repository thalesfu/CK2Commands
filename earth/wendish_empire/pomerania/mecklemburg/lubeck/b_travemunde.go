package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉弗明德TravemundeBarony struct {
	feud.BaseBarony
}

var BTravemunde特拉弗明德 feud.Barony = &特拉弗明德TravemundeBarony{}

func init() {
    f := BTravemunde特拉弗明德.(*特拉弗明德TravemundeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "travemunde",
		TitleName: "特拉弗明德",
		TitleCode: "b_travemunde",
	}
}
