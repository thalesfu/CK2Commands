package bornu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨奥SaoBarony struct {
	feud.BaseBarony
}

var BSao萨奥 feud.Barony = &萨奥SaoBarony{}

func init() {
    f := BSao萨奥.(*萨奥SaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sao",
		TitleName: "萨奥",
		TitleCode: "b_sao",
	}
}
