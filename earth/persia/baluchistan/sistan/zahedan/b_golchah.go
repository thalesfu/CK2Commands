package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉察GolchahBarony struct {
	feud.BaseBarony
}

var BGolchah格拉察 feud.Barony = &格拉察GolchahBarony{}

func init() {
    f := BGolchah格拉察.(*格拉察GolchahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "golchah",
		TitleName: "格拉察",
		TitleCode: "b_golchah",
	}
}
