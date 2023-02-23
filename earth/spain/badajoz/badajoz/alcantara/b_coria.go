package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科里亚CoriaBarony struct {
	feud.BaseBarony
}

var BCoria科里亚 feud.Barony = &科里亚CoriaBarony{}

func init() {
	f := BCoria科里亚.(*科里亚CoriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coria",
		TitleName: "科里亚",
		TitleCode: "b_coria",
	}
}
