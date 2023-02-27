package racakonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗遮军荼RacakondaBarony struct {
	feud.BaseBarony
}

var BRacakonda罗遮军荼 feud.Barony = &罗遮军荼RacakondaBarony{}

func init() {
    f := BRacakonda罗遮军荼.(*罗遮军荼RacakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "racakonda",
		TitleName: "罗遮军荼",
		TitleCode: "b_racakonda",
	}
}
