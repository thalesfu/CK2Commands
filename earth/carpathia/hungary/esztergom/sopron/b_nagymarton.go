package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉毛尔通NagymartonBarony struct {
	feud.BaseBarony
}

var BNagymarton瑙吉毛尔通 feud.Barony = &瑙吉毛尔通NagymartonBarony{}

func init() {
	f := BNagymarton瑙吉毛尔通.(*瑙吉毛尔通NagymartonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagymarton",
		TitleName: "瑙吉毛尔通",
		TitleCode: "b_nagymarton",
	}
}
