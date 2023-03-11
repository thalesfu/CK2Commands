package turov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达维德_戈罗多克Davyd_haradokBarony struct {
	feud.BaseBarony
}

var BDavyd_haradok达维德_戈罗多克 feud.Barony = &达维德_戈罗多克Davyd_haradokBarony{}

func init() {
    f := BDavyd_haradok达维德_戈罗多克.(*达维德_戈罗多克Davyd_haradokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "davyd_haradok",
		TitleName: "达维德_戈罗多克",
		TitleCode: "b_davyd-haradok",
	}
}
