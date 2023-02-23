package kleve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安霍尔特AnholtBarony struct {
	feud.BaseBarony
}

var BAnholt安霍尔特 feud.Barony = &安霍尔特AnholtBarony{}

func init() {
	f := BAnholt安霍尔特.(*安霍尔特AnholtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anholt",
		TitleName: "安霍尔特",
		TitleCode: "b_anholt",
	}
}
