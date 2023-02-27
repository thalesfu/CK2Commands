package kumul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大河DaheBarony struct {
	feud.BaseBarony
}

var BDahe大河 feud.Barony = &大河DaheBarony{}

func init() {
    f := BDahe大河.(*大河DaheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dahe",
		TitleName: "大河",
		TitleCode: "b_dahe",
	}
}
