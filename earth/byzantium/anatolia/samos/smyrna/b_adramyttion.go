package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德拉密提翁AdramyttionBarony struct {
	feud.BaseBarony
}

var BAdramyttion阿德拉密提翁 feud.Barony = &阿德拉密提翁AdramyttionBarony{}

func init() {
    f := BAdramyttion阿德拉密提翁.(*阿德拉密提翁AdramyttionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adramyttion",
		TitleName: "阿德拉密提翁",
		TitleCode: "b_adramyttion",
	}
}
