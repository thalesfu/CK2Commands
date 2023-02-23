package kollipake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗底耶RatiaBarony struct {
	feud.BaseBarony
}

var BRatia罗底耶 feud.Barony = &罗底耶RatiaBarony{}

func init() {
	f := BRatia罗底耶.(*罗底耶RatiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ratia",
		TitleName: "罗底耶",
		TitleCode: "b_ratia",
	}
}
