package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特塞尔TexelBarony struct {
	feud.BaseBarony
}

var BTexel特塞尔 feud.Barony = &特塞尔TexelBarony{}

func init() {
	f := BTexel特塞尔.(*特塞尔TexelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "texel",
		TitleName: "特塞尔",
		TitleCode: "b_texel",
	}
}
