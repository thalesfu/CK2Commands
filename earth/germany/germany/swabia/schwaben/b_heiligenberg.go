package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海利根贝格HeiligenbergBarony struct {
	feud.BaseBarony
}

var BHeiligenberg海利根贝格 feud.Barony = &海利根贝格HeiligenbergBarony{}

func init() {
	f := BHeiligenberg海利根贝格.(*海利根贝格HeiligenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heiligenberg",
		TitleName: "海利根贝格",
		TitleCode: "b_heiligenberg",
	}
}
