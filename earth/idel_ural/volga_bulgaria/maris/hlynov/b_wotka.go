package hlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃特卡WotkaBarony struct {
	feud.BaseBarony
}

var BWotka沃特卡 feud.Barony = &沃特卡WotkaBarony{}

func init() {
    f := BWotka沃特卡.(*沃特卡WotkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wotka",
		TitleName: "沃特卡",
		TitleCode: "b_wotka",
	}
}
