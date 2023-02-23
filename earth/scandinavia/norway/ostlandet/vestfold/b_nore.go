package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努勒NoreBarony struct {
	feud.BaseBarony
}

var BNore努勒 feud.Barony = &努勒NoreBarony{}

func init() {
	f := BNore努勒.(*努勒NoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nore",
		TitleName: "努勒",
		TitleCode: "b_nore",
	}
}
