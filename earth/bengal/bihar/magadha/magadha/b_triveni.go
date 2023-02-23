package magadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝力吠尼TriveniBarony struct {
	feud.BaseBarony
}

var BTriveni帝力吠尼 feud.Barony = &帝力吠尼TriveniBarony{}

func init() {
	f := BTriveni帝力吠尼.(*帝力吠尼TriveniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "triveni",
		TitleName: "帝力吠尼",
		TitleCode: "b_triveni",
	}
}
