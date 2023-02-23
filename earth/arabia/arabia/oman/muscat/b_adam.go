package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚当AdamBarony struct {
	feud.BaseBarony
}

var BAdam亚当 feud.Barony = &亚当AdamBarony{}

func init() {
	f := BAdam亚当.(*亚当AdamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adam",
		TitleName: "亚当",
		TitleCode: "b_adam",
	}
}
