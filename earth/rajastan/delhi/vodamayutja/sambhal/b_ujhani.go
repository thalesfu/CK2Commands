package sambhal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌诃尼UjhaniBarony struct {
	feud.BaseBarony
}

var BUjhani乌诃尼 feud.Barony = &乌诃尼UjhaniBarony{}

func init() {
	f := BUjhani乌诃尼.(*乌诃尼UjhaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ujhani",
		TitleName: "乌诃尼",
		TitleCode: "b_ujhani",
	}
}
