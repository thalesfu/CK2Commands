package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利马索尔LimmassolBarony struct {
	feud.BaseBarony
}

var BLimmassol利马索尔 feud.Barony = &利马索尔LimmassolBarony{}

func init() {
	f := BLimmassol利马索尔.(*利马索尔LimmassolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "limmassol",
		TitleName: "利马索尔",
		TitleCode: "b_limmassol",
	}
}
