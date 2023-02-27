package kazakh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎伊列姆ZhairemBarony struct {
	feud.BaseBarony
}

var BZhairem扎伊列姆 feud.Barony = &扎伊列姆ZhairemBarony{}

func init() {
    f := BZhairem扎伊列姆.(*扎伊列姆ZhairemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhairem",
		TitleName: "扎伊列姆",
		TitleCode: "b_zhairem",
	}
}
