package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔米耶UrmiahBarony struct {
	feud.BaseBarony
}

var BUrmiah乌尔米耶 feud.Barony = &乌尔米耶UrmiahBarony{}

func init() {
	f := BUrmiah乌尔米耶.(*乌尔米耶UrmiahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urmiah",
		TitleName: "乌尔米耶",
		TitleCode: "b_urmiah",
	}
}
