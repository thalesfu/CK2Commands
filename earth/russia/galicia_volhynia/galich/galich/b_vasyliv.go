package galich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦西列夫VasylivBarony struct {
	feud.BaseBarony
}

var BVasyliv瓦西列夫 feud.Barony = &瓦西列夫VasylivBarony{}

func init() {
    f := BVasyliv瓦西列夫.(*瓦西列夫VasylivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vasyliv",
		TitleName: "瓦西列夫",
		TitleCode: "b_vasyliv",
	}
}
