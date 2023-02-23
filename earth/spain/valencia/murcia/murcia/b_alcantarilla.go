package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔坎塔里利亚AlcantarillaBarony struct {
	feud.BaseBarony
}

var BAlcantarilla阿尔坎塔里利亚 feud.Barony = &阿尔坎塔里利亚AlcantarillaBarony{}

func init() {
	f := BAlcantarilla阿尔坎塔里利亚.(*阿尔坎塔里利亚AlcantarillaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcantarilla",
		TitleName: "阿尔坎塔里利亚",
		TitleCode: "b_alcantarilla",
	}
}
