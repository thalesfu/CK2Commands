package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因拜尔德AinelberdBarony struct {
	feud.BaseBarony
}

var BAinelberd艾因拜尔德 feud.Barony = &艾因拜尔德AinelberdBarony{}

func init() {
	f := BAinelberd艾因拜尔德.(*艾因拜尔德AinelberdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainelberd",
		TitleName: "艾因拜尔德",
		TitleCode: "b_ainelberd",
	}
}
