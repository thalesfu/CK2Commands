package qaidam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡铁山XitishanBarony struct {
	feud.BaseBarony
}

var BXitishan锡铁山 feud.Barony = &锡铁山XitishanBarony{}

func init() {
    f := BXitishan锡铁山.(*锡铁山XitishanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xitishan",
		TitleName: "锡铁山",
		TitleCode: "b_xitishan",
	}
}
