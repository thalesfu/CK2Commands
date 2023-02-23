package chlisselbourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯托罗日诺StorozhnoBarony struct {
	feud.BaseBarony
}

var BStorozhno斯托罗日诺 feud.Barony = &斯托罗日诺StorozhnoBarony{}

func init() {
	f := BStorozhno斯托罗日诺.(*斯托罗日诺StorozhnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "storozhno",
		TitleName: "斯托罗日诺",
		TitleCode: "b_storozhno",
	}
}
