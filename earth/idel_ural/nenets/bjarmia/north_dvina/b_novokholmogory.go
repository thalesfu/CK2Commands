package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新霍尔莫戈雷NovokholmogoryBarony struct {
	feud.BaseBarony
}

var BNovokholmogory新霍尔莫戈雷 feud.Barony = &新霍尔莫戈雷NovokholmogoryBarony{}

func init() {
    f := BNovokholmogory新霍尔莫戈雷.(*新霍尔莫戈雷NovokholmogoryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novokholmogory",
		TitleName: "新霍尔莫戈雷",
		TitleCode: "b_novokholmogory",
	}
}
