package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃吉耶拉VoghieraBarony struct {
	feud.BaseBarony
}

var BVoghiera沃吉耶拉 feud.Barony = &沃吉耶拉VoghieraBarony{}

func init() {
    f := BVoghiera沃吉耶拉.(*沃吉耶拉VoghieraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voghiera",
		TitleName: "沃吉耶拉",
		TitleCode: "b_voghiera",
	}
}
