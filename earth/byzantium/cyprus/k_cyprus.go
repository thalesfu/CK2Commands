package cyprus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CyprusKingdom interface {
	feud.Kingdom
}

type 塞浦路斯CyprusKingdom struct {
	feud.BaseKingdom
}

var KCyprus塞浦路斯 CyprusKingdom = &塞浦路斯CyprusKingdom{}

func init() {
	f := KCyprus塞浦路斯.(*塞浦路斯CyprusKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "cyprus",
		TitleName: "塞浦路斯",
		TitleCode: "k_cyprus",
		Dukes:     map[string]feud.Duke{},
	}

}
