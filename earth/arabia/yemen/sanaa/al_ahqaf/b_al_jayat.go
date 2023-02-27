package al_ahqaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰阿特Al_jayatBarony struct {
	feud.BaseBarony
}

var BAl_jayat杰阿特 feud.Barony = &杰阿特Al_jayatBarony{}

func init() {
    f := BAl_jayat杰阿特.(*杰阿特Al_jayatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_jayat",
		TitleName: "杰阿特",
		TitleCode: "b_al_jayat",
	}
}
