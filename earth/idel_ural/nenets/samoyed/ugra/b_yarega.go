package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚列加YaregaBarony struct {
	feud.BaseBarony
}

var BYarega亚列加 feud.Barony = &亚列加YaregaBarony{}

func init() {
    f := BYarega亚列加.(*亚列加YaregaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yarega",
		TitleName: "亚列加",
		TitleCode: "b_yarega",
	}
}
