package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡吉勒马萨SijilmasaBarony struct {
	feud.BaseBarony
}

var BSijilmasa锡吉勒马萨 feud.Barony = &锡吉勒马萨SijilmasaBarony{}

func init() {
	f := BSijilmasa锡吉勒马萨.(*锡吉勒马萨SijilmasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sijilmasa",
		TitleName: "锡吉勒马萨",
		TitleCode: "b_sijilmasa",
	}
}
