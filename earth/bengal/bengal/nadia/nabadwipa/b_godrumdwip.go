package nabadwipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿度摩洲GodrumdwipBarony struct {
	feud.BaseBarony
}

var BGodrumdwip瞿度摩洲 feud.Barony = &瞿度摩洲GodrumdwipBarony{}

func init() {
    f := BGodrumdwip瞿度摩洲.(*瞿度摩洲GodrumdwipBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "godrumdwip",
		TitleName: "瞿度摩洲",
		TitleCode: "b_godrumdwip",
	}
}
