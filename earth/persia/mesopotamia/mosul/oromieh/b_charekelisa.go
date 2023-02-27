package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察雷克里撒CharekelisaBarony struct {
	feud.BaseBarony
}

var BCharekelisa察雷克里撒 feud.Barony = &察雷克里撒CharekelisaBarony{}

func init() {
    f := BCharekelisa察雷克里撒.(*察雷克里撒CharekelisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charekelisa",
		TitleName: "察雷克里撒",
		TitleCode: "b_charekelisa",
	}
}
