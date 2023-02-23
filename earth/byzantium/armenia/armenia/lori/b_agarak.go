package lori

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿加拉克AgarakBarony struct {
	feud.BaseBarony
}

var BAgarak阿加拉克 feud.Barony = &阿加拉克AgarakBarony{}

func init() {
	f := BAgarak阿加拉克.(*阿加拉克AgarakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agarak",
		TitleName: "阿加拉克",
		TitleCode: "b_agarak",
	}
}
