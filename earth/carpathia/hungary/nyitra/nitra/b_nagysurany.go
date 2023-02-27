package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉舒拉尼NagysuranyBarony struct {
	feud.BaseBarony
}

var BNagysurany瑙吉舒拉尼 feud.Barony = &瑙吉舒拉尼NagysuranyBarony{}

func init() {
    f := BNagysurany瑙吉舒拉尼.(*瑙吉舒拉尼NagysuranyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagysurany",
		TitleName: "瑙吉舒拉尼",
		TitleCode: "b_nagysurany",
	}
}
