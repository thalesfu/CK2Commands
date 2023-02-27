package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈登HawardenBarony struct {
	feud.BaseBarony
}

var BHawarden哈登 feud.Barony = &哈登HawardenBarony{}

func init() {
    f := BHawarden哈登.(*哈登HawardenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hawarden",
		TitleName: "哈登",
		TitleCode: "b_hawarden",
	}
}
