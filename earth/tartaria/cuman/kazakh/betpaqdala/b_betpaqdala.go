package betpaqdala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别特帕克达拉BetpaqdalaBarony struct {
	feud.BaseBarony
}

var BBetpaqdala别特帕克达拉 feud.Barony = &别特帕克达拉BetpaqdalaBarony{}

func init() {
    f := BBetpaqdala别特帕克达拉.(*别特帕克达拉BetpaqdalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "betpaqdala",
		TitleName: "别特帕克达拉",
		TitleCode: "b_betpaqdala",
	}
}
