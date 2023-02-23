package bayuda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉代格ElkhandaqBarony struct {
	feud.BaseBarony
}

var BElkhandaq汉代格 feud.Barony = &汉代格ElkhandaqBarony{}

func init() {
	f := BElkhandaq汉代格.(*汉代格ElkhandaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elkhandaq",
		TitleName: "汉代格",
		TitleCode: "b_elkhandaq",
	}
}
