package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特埃莫索MontehermosoBarony struct {
	feud.BaseBarony
}

var BMontehermoso蒙特埃莫索 feud.Barony = &蒙特埃莫索MontehermosoBarony{}

func init() {
	f := BMontehermoso蒙特埃莫索.(*蒙特埃莫索MontehermosoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montehermoso",
		TitleName: "蒙特埃莫索",
		TitleCode: "b_montehermoso",
	}
}
