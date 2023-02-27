package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔齐尼亚诺ArzignanoBarony struct {
	feud.BaseBarony
}

var BArzignano阿尔齐尼亚诺 feud.Barony = &阿尔齐尼亚诺ArzignanoBarony{}

func init() {
    f := BArzignano阿尔齐尼亚诺.(*阿尔齐尼亚诺ArzignanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arzignano",
		TitleName: "阿尔齐尼亚诺",
		TitleCode: "b_arzignano",
	}
}
