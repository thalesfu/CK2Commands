package bath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯纳姆BurnhamBarony struct {
	feud.BaseBarony
}

var BBurnham伯纳姆 feud.Barony = &伯纳姆BurnhamBarony{}

func init() {
	f := BBurnham伯纳姆.(*伯纳姆BurnhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burnham",
		TitleName: "伯纳姆",
		TitleCode: "b_burnham",
	}
}
