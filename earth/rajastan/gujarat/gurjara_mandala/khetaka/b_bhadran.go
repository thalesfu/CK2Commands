package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波多兰BhadranBarony struct {
	feud.BaseBarony
}

var BBhadran波多兰 feud.Barony = &波多兰BhadranBarony{}

func init() {
    f := BBhadran波多兰.(*波多兰BhadranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadran",
		TitleName: "波多兰",
		TitleCode: "b_bhadran",
	}
}
