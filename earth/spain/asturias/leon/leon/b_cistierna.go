package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡斯铁尔纳CistiernaBarony struct {
	feud.BaseBarony
}

var BCistierna锡斯铁尔纳 feud.Barony = &锡斯铁尔纳CistiernaBarony{}

func init() {
	f := BCistierna锡斯铁尔纳.(*锡斯铁尔纳CistiernaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cistierna",
		TitleName: "锡斯铁尔纳",
		TitleCode: "b_cistierna",
	}
}
