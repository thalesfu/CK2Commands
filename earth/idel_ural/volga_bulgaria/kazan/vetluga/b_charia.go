package vetluga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙里亚ChariaBarony struct {
	feud.BaseBarony
}

var BCharia沙里亚 feud.Barony = &沙里亚ChariaBarony{}

func init() {
    f := BCharia沙里亚.(*沙里亚ChariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charia",
		TitleName: "沙里亚",
		TitleCode: "b_charia",
	}
}
