package consenza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗萨诺RossanoBarony struct {
	feud.BaseBarony
}

var BRossano罗萨诺 feud.Barony = &罗萨诺RossanoBarony{}

func init() {
	f := BRossano罗萨诺.(*罗萨诺RossanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rossano",
		TitleName: "罗萨诺",
		TitleCode: "b_rossano",
	}
}
