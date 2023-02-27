package hull

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯坦福桥Stamford_bridgeBarony struct {
	feud.BaseBarony
}

var BStamford_bridge斯坦福桥 feud.Barony = &斯坦福桥Stamford_bridgeBarony{}

func init() {
    f := BStamford_bridge斯坦福桥.(*斯坦福桥Stamford_bridgeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stamford_bridge",
		TitleName: "斯坦福桥",
		TitleCode: "b_stamford_bridge",
	}
}
