package kherla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格维尔格尔GawilgharBarony struct {
	feud.BaseBarony
}

var BGawilghar格维尔格尔 feud.Barony = &格维尔格尔GawilgharBarony{}

func init() {
    f := BGawilghar格维尔格尔.(*格维尔格尔GawilgharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gawilghar",
		TitleName: "格维尔格尔",
		TitleCode: "b_gawilghar",
	}
}
