package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布吉尔AbukirBarony struct {
	feud.BaseBarony
}

var BAbukir阿布吉尔 feud.Barony = &阿布吉尔AbukirBarony{}

func init() {
    f := BAbukir阿布吉尔.(*阿布吉尔AbukirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abukir",
		TitleName: "阿布吉尔",
		TitleCode: "b_abukir",
	}
}
