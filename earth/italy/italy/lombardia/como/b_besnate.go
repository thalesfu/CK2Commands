package como

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝斯纳泰BesnateBarony struct {
	feud.BaseBarony
}

var BBesnate贝斯纳泰 feud.Barony = &贝斯纳泰BesnateBarony{}

func init() {
    f := BBesnate贝斯纳泰.(*贝斯纳泰BesnateBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "besnate",
		TitleName: "贝斯纳泰",
		TitleCode: "b_besnate",
	}
}
