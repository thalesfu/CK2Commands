package mingoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡克沁ShikshinBarony struct {
	feud.BaseBarony
}

var BShikshin锡克沁 feud.Barony = &锡克沁ShikshinBarony{}

func init() {
	f := BShikshin锡克沁.(*锡克沁ShikshinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shikshin",
		TitleName: "锡克沁",
		TitleCode: "b_shikshin",
	}
}
