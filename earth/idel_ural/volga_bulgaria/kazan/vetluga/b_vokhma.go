package vetluga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃赫马VokhmaBarony struct {
	feud.BaseBarony
}

var BVokhma沃赫马 feud.Barony = &沃赫马VokhmaBarony{}

func init() {
    f := BVokhma沃赫马.(*沃赫马VokhmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vokhma",
		TitleName: "沃赫马",
		TitleCode: "b_vokhma",
	}
}
