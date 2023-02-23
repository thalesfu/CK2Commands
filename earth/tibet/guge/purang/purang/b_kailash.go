package purang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸡罗娑KailashBarony struct {
	feud.BaseBarony
}

var BKailash鸡罗娑 feud.Barony = &鸡罗娑KailashBarony{}

func init() {
	f := BKailash鸡罗娑.(*鸡罗娑KailashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kailash",
		TitleName: "鸡罗娑",
		TitleCode: "b_kailash",
	}
}
