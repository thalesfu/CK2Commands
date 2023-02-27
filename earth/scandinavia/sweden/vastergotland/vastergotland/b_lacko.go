package vastergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱雪LackoBarony struct {
	feud.BaseBarony
}

var BLacko莱雪 feud.Barony = &莱雪LackoBarony{}

func init() {
    f := BLacko莱雪.(*莱雪LackoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lacko",
		TitleName: "莱雪",
		TitleCode: "b_lacko",
	}
}
