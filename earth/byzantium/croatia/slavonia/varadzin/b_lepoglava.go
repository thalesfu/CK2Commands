package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱波格拉瓦LepoglavaBarony struct {
	feud.BaseBarony
}

var BLepoglava莱波格拉瓦 feud.Barony = &莱波格拉瓦LepoglavaBarony{}

func init() {
	f := BLepoglava莱波格拉瓦.(*莱波格拉瓦LepoglavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lepoglava",
		TitleName: "莱波格拉瓦",
		TitleCode: "b_lepoglava",
	}
}
