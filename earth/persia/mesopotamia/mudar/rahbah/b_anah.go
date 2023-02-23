package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿纳海AnahBarony struct {
	feud.BaseBarony
}

var BAnah阿纳海 feud.Barony = &阿纳海AnahBarony{}

func init() {
	f := BAnah阿纳海.(*阿纳海AnahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anah",
		TitleName: "阿纳海",
		TitleCode: "b_anah",
	}
}
