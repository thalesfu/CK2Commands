package somerset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉斯顿伯里GlastonburyBarony struct {
	feud.BaseBarony
}

var BGlastonbury格拉斯顿伯里 feud.Barony = &格拉斯顿伯里GlastonburyBarony{}

func init() {
	f := BGlastonbury格拉斯顿伯里.(*格拉斯顿伯里GlastonburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glastonbury",
		TitleName: "格拉斯顿伯里",
		TitleCode: "b_glastonbury",
	}
}
