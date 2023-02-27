package sambalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗阇恒伽补罗RajgangpurBarony struct {
	feud.BaseBarony
}

var BRajgangpur罗阇恒伽补罗 feud.Barony = &罗阇恒伽补罗RajgangpurBarony{}

func init() {
    f := BRajgangpur罗阇恒伽补罗.(*罗阇恒伽补罗RajgangpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajgangpur",
		TitleName: "罗阇恒伽补罗",
		TitleCode: "b_rajgangpur",
	}
}
