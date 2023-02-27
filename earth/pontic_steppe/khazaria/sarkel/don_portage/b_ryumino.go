package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 留米诺RyuminoBarony struct {
	feud.BaseBarony
}

var BRyumino留米诺 feud.Barony = &留米诺RyuminoBarony{}

func init() {
    f := BRyumino留米诺.(*留米诺RyuminoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ryumino",
		TitleName: "留米诺",
		TitleCode: "b_ryumino",
	}
}
