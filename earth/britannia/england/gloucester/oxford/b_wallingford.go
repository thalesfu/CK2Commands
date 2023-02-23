package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃灵福德WallingfordBarony struct {
	feud.BaseBarony
}

var BWallingford沃灵福德 feud.Barony = &沃灵福德WallingfordBarony{}

func init() {
	f := BWallingford沃灵福德.(*沃灵福德WallingfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wallingford",
		TitleName: "沃灵福德",
		TitleCode: "b_wallingford",
	}
}
