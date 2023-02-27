package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉比亚At_talabiyaBarony struct {
	feud.BaseBarony
}

var BAt_talabiya塔拉比亚 feud.Barony = &塔拉比亚At_talabiyaBarony{}

func init() {
    f := BAt_talabiya塔拉比亚.(*塔拉比亚At_talabiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "at_talabiya",
		TitleName: "塔拉比亚",
		TitleCode: "b_at_talabiya",
	}
}
