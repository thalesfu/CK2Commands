package ivrea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尚波尔谢ChamporcherBarony struct {
	feud.BaseBarony
}

var BChamporcher尚波尔谢 feud.Barony = &尚波尔谢ChamporcherBarony{}

func init() {
	f := BChamporcher尚波尔谢.(*尚波尔谢ChamporcherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "champorcher",
		TitleName: "尚波尔谢",
		TitleCode: "b_champorcher",
	}
}
