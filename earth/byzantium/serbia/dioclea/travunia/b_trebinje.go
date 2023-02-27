package travunia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷比涅TrebinjeBarony struct {
	feud.BaseBarony
}

var BTrebinje特雷比涅 feud.Barony = &特雷比涅TrebinjeBarony{}

func init() {
    f := BTrebinje特雷比涅.(*特雷比涅TrebinjeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trebinje",
		TitleName: "特雷比涅",
		TitleCode: "b_trebinje",
	}
}
