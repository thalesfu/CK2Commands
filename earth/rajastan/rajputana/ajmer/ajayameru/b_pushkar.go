package ajayameru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布色羯逻PushkarBarony struct {
	feud.BaseBarony
}

var BPushkar布色羯逻 feud.Barony = &布色羯逻PushkarBarony{}

func init() {
    f := BPushkar布色羯逻.(*布色羯逻PushkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pushkar",
		TitleName: "布色羯逻",
		TitleCode: "b_pushkar",
	}
}
