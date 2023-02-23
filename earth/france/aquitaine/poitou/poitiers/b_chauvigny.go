package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍维尼ChauvignyBarony struct {
	feud.BaseBarony
}

var BChauvigny绍维尼 feud.Barony = &绍维尼ChauvignyBarony{}

func init() {
	f := BChauvigny绍维尼.(*绍维尼ChauvignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chauvigny",
		TitleName: "绍维尼",
		TitleCode: "b_chauvigny",
	}
}
