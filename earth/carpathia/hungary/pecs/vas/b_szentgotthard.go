package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣戈特哈德SzentgotthardBarony struct {
	feud.BaseBarony
}

var BSzentgotthard圣戈特哈德 feud.Barony = &圣戈特哈德SzentgotthardBarony{}

func init() {
	f := BSzentgotthard圣戈特哈德.(*圣戈特哈德SzentgotthardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szentgotthard",
		TitleName: "圣戈特哈德",
		TitleCode: "b_szentgotthard",
	}
}
