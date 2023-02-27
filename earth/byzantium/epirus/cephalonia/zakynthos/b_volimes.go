package zakynthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃利梅斯VolimesBarony struct {
	feud.BaseBarony
}

var BVolimes沃利梅斯 feud.Barony = &沃利梅斯VolimesBarony{}

func init() {
    f := BVolimes沃利梅斯.(*沃利梅斯VolimesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volimes",
		TitleName: "沃利梅斯",
		TitleCode: "b_volimes",
	}
}
