package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海法HaifaBarony struct {
	feud.BaseBarony
}

var BHaifa海法 feud.Barony = &海法HaifaBarony{}

func init() {
	f := BHaifa海法.(*海法HaifaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haifa",
		TitleName: "海法",
		TitleCode: "b_haifa",
	}
}
