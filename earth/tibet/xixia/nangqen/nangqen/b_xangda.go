package nangqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 香达XangdaBarony struct {
	feud.BaseBarony
}

var BXangda香达 feud.Barony = &香达XangdaBarony{}

func init() {
    f := BXangda香达.(*香达XangdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xangda",
		TitleName: "香达",
		TitleCode: "b_xangda",
	}
}
