package kirghiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 铁勒TeleBarony struct {
	feud.BaseBarony
}

var BTele铁勒 feud.Barony = &铁勒TeleBarony{}

func init() {
	f := BTele铁勒.(*铁勒TeleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tele",
		TitleName: "铁勒",
		TitleCode: "b_tele",
	}
}
