package torzhok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉什基诺RashkinoBarony struct {
	feud.BaseBarony
}

var BRashkino拉什基诺 feud.Barony = &拉什基诺RashkinoBarony{}

func init() {
    f := BRashkino拉什基诺.(*拉什基诺RashkinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rashkino",
		TitleName: "拉什基诺",
		TitleCode: "b_rashkino",
	}
}
