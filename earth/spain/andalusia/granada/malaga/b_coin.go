package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科因CoinBarony struct {
	feud.BaseBarony
}

var BCoin科因 feud.Barony = &科因CoinBarony{}

func init() {
    f := BCoin科因.(*科因CoinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coin",
		TitleName: "科因",
		TitleCode: "b_coin",
	}
}
