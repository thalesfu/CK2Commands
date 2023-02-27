package kudalasangama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敦巴尔DambalBarony struct {
	feud.BaseBarony
}

var BDambal敦巴尔 feud.Barony = &敦巴尔DambalBarony{}

func init() {
    f := BDambal敦巴尔.(*敦巴尔DambalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dambal",
		TitleName: "敦巴尔",
		TitleCode: "b_dambal",
	}
}
