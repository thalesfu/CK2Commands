package jharkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菴毗迦城AmbikapurBarony struct {
	feud.BaseBarony
}

var BAmbikapur菴毗迦城 feud.Barony = &菴毗迦城AmbikapurBarony{}

func init() {
	f := BAmbikapur菴毗迦城.(*菴毗迦城AmbikapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ambikapur",
		TitleName: "菴毗迦城",
		TitleCode: "b_ambikapur",
	}
}
