package nanded

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔博亚尼ParbhaniBarony struct {
	feud.BaseBarony
}

var BParbhani帕尔博亚尼 feud.Barony = &帕尔博亚尼ParbhaniBarony{}

func init() {
    f := BParbhani帕尔博亚尼.(*帕尔博亚尼ParbhaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parbhani",
		TitleName: "帕尔博亚尼",
		TitleCode: "b_parbhani",
	}
}
