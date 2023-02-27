package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔略CharlieuBarony struct {
	feud.BaseBarony
}

var BCharlieu沙尔略 feud.Barony = &沙尔略CharlieuBarony{}

func init() {
    f := BCharlieu沙尔略.(*沙尔略CharlieuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charlieu",
		TitleName: "沙尔略",
		TitleCode: "b_charlieu",
	}
}
