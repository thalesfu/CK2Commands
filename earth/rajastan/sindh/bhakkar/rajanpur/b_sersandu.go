package rajanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟桑堵SersanduBarony struct {
	feud.BaseBarony
}

var BSersandu瑟桑堵 feud.Barony = &瑟桑堵SersanduBarony{}

func init() {
    f := BSersandu瑟桑堵.(*瑟桑堵SersanduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sersandu",
		TitleName: "瑟桑堵",
		TitleCode: "b_sersandu",
	}
}
