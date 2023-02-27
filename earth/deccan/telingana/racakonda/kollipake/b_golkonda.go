package kollipake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿罗军荼GolkondaBarony struct {
	feud.BaseBarony
}

var BGolkonda瞿罗军荼 feud.Barony = &瞿罗军荼GolkondaBarony{}

func init() {
    f := BGolkonda瞿罗军荼.(*瞿罗军荼GolkondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "golkonda",
		TitleName: "瞿罗军荼",
		TitleCode: "b_golkonda",
	}
}
