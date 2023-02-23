package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦利亚VegliaBarony struct {
	feud.BaseBarony
}

var BVeglia韦利亚 feud.Barony = &韦利亚VegliaBarony{}

func init() {
	f := BVeglia韦利亚.(*韦利亚VegliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veglia",
		TitleName: "韦利亚",
		TitleCode: "b_veglia",
	}
}
