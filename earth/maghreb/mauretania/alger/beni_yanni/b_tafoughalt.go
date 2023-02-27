package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔夫亚拉特TafoughaltBarony struct {
	feud.BaseBarony
}

var BTafoughalt塔夫亚拉特 feud.Barony = &塔夫亚拉特TafoughaltBarony{}

func init() {
    f := BTafoughalt塔夫亚拉特.(*塔夫亚拉特TafoughaltBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tafoughalt",
		TitleName: "塔夫亚拉特",
		TitleCode: "b_tafoughalt",
	}
}
