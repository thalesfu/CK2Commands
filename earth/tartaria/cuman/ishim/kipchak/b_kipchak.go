package kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钦察KipchakBarony struct {
	feud.BaseBarony
}

var BKipchak钦察 feud.Barony = &钦察KipchakBarony{}

func init() {
    f := BKipchak钦察.(*钦察KipchakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kipchak",
		TitleName: "钦察",
		TitleCode: "b_kipchak",
	}
}
