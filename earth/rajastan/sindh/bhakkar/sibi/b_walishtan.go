package sibi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦利斯坦WalishtanBarony struct {
	feud.BaseBarony
}

var BWalishtan瓦利斯坦 feud.Barony = &瓦利斯坦WalishtanBarony{}

func init() {
    f := BWalishtan瓦利斯坦.(*瓦利斯坦WalishtanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "walishtan",
		TitleName: "瓦利斯坦",
		TitleCode: "b_walishtan",
	}
}
