package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔杜VerduBarony struct {
	feud.BaseBarony
}

var BVerdu贝尔杜 feud.Barony = &贝尔杜VerduBarony{}

func init() {
	f := BVerdu贝尔杜.(*贝尔杜VerduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verdu",
		TitleName: "贝尔杜",
		TitleCode: "b_verdu",
	}
}
