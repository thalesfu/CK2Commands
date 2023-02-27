package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗咩罗旬迦罗Bhimer_jangalBarony struct {
	feud.BaseBarony
}

var BBhimer_jangal毗咩罗旬迦罗 feud.Barony = &毗咩罗旬迦罗Bhimer_jangalBarony{}

func init() {
    f := BBhimer_jangal毗咩罗旬迦罗.(*毗咩罗旬迦罗Bhimer_jangalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhimer_jangal",
		TitleName: "毗咩罗旬迦罗",
		TitleCode: "b_bhimer_jangal",
	}
}
