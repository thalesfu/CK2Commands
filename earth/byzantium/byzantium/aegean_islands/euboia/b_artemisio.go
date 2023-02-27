package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿提米修ArtemisioBarony struct {
	feud.BaseBarony
}

var BArtemisio阿提米修 feud.Barony = &阿提米修ArtemisioBarony{}

func init() {
    f := BArtemisio阿提米修.(*阿提米修ArtemisioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "artemisio",
		TitleName: "阿提米修",
		TitleCode: "b_artemisio",
	}
}
