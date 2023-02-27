package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡塞巴HusaibaBarony struct {
	feud.BaseBarony
}

var BHusaiba胡塞巴 feud.Barony = &胡塞巴HusaibaBarony{}

func init() {
    f := BHusaiba胡塞巴.(*胡塞巴HusaibaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "husaiba",
		TitleName: "胡塞巴",
		TitleCode: "b_husaiba",
	}
}
