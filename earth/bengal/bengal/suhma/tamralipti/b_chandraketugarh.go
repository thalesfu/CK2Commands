package tamralipti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀罗计都姞利呬ChandraketugarhBarony struct {
	feud.BaseBarony
}

var BChandraketugarh旃陀罗计都姞利呬 feud.Barony = &旃陀罗计都姞利呬ChandraketugarhBarony{}

func init() {
    f := BChandraketugarh旃陀罗计都姞利呬.(*旃陀罗计都姞利呬ChandraketugarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chandraketugarh",
		TitleName: "旃陀罗计都姞利呬",
		TitleCode: "b_chandraketugarh",
	}
}
