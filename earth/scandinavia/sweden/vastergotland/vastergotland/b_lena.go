package vastergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱纳LenaBarony struct {
	feud.BaseBarony
}

var BLena莱纳 feud.Barony = &莱纳LenaBarony{}

func init() {
    f := BLena莱纳.(*莱纳LenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lena",
		TitleName: "莱纳",
		TitleCode: "b_lena",
	}
}
