package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡尔萨KhursaBarony struct {
	feud.BaseBarony
}

var BKhursa胡尔萨 feud.Barony = &胡尔萨KhursaBarony{}

func init() {
    f := BKhursa胡尔萨.(*胡尔萨KhursaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khursa",
		TitleName: "胡尔萨",
		TitleCode: "b_khursa",
	}
}
