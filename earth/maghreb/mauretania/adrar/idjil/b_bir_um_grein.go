package idjil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔乌姆格兰Bir_um_greinBarony struct {
	feud.BaseBarony
}

var BBir_um_grein比尔乌姆格兰 feud.Barony = &比尔乌姆格兰Bir_um_greinBarony{}

func init() {
    f := BBir_um_grein比尔乌姆格兰.(*比尔乌姆格兰Bir_um_greinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bir_um_grein",
		TitleName: "比尔乌姆格兰",
		TitleCode: "b_bir_um_grein",
	}
}
