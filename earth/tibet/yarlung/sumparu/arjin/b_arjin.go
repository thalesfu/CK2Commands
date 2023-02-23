package arjin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔金ArjinBarony struct {
	feud.BaseBarony
}

var BArjin阿尔金 feud.Barony = &阿尔金ArjinBarony{}

func init() {
	f := BArjin阿尔金.(*阿尔金ArjinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arjin",
		TitleName: "阿尔金",
		TitleCode: "b_arjin",
	}
}
