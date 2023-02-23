package philippopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 腓利波波利斯PhilippopolisBarony struct {
	feud.BaseBarony
}

var BPhilippopolis腓利波波利斯 feud.Barony = &腓利波波利斯PhilippopolisBarony{}

func init() {
	f := BPhilippopolis腓利波波利斯.(*腓利波波利斯PhilippopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "philippopolis",
		TitleName: "腓利波波利斯",
		TitleCode: "b_philippopolis",
	}
}
