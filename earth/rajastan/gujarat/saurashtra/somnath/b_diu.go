package somnath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 刁元DiuBarony struct {
	feud.BaseBarony
}

var BDiu刁元 feud.Barony = &刁元DiuBarony{}

func init() {
    f := BDiu刁元.(*刁元DiuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diu",
		TitleName: "刁元",
		TitleCode: "b_diu",
	}
}
