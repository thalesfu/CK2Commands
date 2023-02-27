package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉德梅尔斯基GradmerskyBarony struct {
	feud.BaseBarony
}

var BGradmersky格拉德梅尔斯基 feud.Barony = &格拉德梅尔斯基GradmerskyBarony{}

func init() {
    f := BGradmersky格拉德梅尔斯基.(*格拉德梅尔斯基GradmerskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gradmersky",
		TitleName: "格拉德梅尔斯基",
		TitleCode: "b_gradmersky",
	}
}
