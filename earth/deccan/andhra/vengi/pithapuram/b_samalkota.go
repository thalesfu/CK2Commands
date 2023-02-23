package pithapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨马尔科塔SamalkotaBarony struct {
	feud.BaseBarony
}

var BSamalkota萨马尔科塔 feud.Barony = &萨马尔科塔SamalkotaBarony{}

func init() {
	f := BSamalkota萨马尔科塔.(*萨马尔科塔SamalkotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samalkota",
		TitleName: "萨马尔科塔",
		TitleCode: "b_samalkota",
	}
}
