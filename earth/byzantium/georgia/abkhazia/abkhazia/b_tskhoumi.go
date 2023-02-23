package abkhazia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨胡米TskhoumiBarony struct {
	feud.BaseBarony
}

var BTskhoumi茨胡米 feud.Barony = &茨胡米TskhoumiBarony{}

func init() {
	f := BTskhoumi茨胡米.(*茨胡米TskhoumiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tskhoumi",
		TitleName: "茨胡米",
		TitleCode: "b_tskhoumi",
	}
}
