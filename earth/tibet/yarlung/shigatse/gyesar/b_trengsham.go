package gyesar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 惩香TrengshamBarony struct {
	feud.BaseBarony
}

var BTrengsham惩香 feud.Barony = &惩香TrengshamBarony{}

func init() {
	f := BTrengsham惩香.(*惩香TrengshamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trengsham",
		TitleName: "惩香",
		TitleCode: "b_trengsham",
	}
}
