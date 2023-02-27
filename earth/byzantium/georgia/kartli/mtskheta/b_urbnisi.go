package mtskheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔布尼西UrbnisiBarony struct {
	feud.BaseBarony
}

var BUrbnisi乌尔布尼西 feud.Barony = &乌尔布尼西UrbnisiBarony{}

func init() {
    f := BUrbnisi乌尔布尼西.(*乌尔布尼西UrbnisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urbnisi",
		TitleName: "乌尔布尼西",
		TitleCode: "b_urbnisi",
	}
}
