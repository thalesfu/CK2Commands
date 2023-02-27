package ural

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克什特姆KyshtymBarony struct {
	feud.BaseBarony
}

var BKyshtym克什特姆 feud.Barony = &克什特姆KyshtymBarony{}

func init() {
    f := BKyshtym克什特姆.(*克什特姆KyshtymBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyshtym",
		TitleName: "克什特姆",
		TitleCode: "b_kyshtym",
	}
}
