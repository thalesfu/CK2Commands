package aylik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔禾UrhoBarony struct {
	feud.BaseBarony
}

var BUrho乌尔禾 feud.Barony = &乌尔禾UrhoBarony{}

func init() {
    f := BUrho乌尔禾.(*乌尔禾UrhoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urho",
		TitleName: "乌尔禾",
		TitleCode: "b_urho",
	}
}
