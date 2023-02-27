package czersk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布鲁德诺BrodnoBarony struct {
	feud.BaseBarony
}

var BBrodno布鲁德诺 feud.Barony = &布鲁德诺BrodnoBarony{}

func init() {
    f := BBrodno布鲁德诺.(*布鲁德诺BrodnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brodno",
		TitleName: "布鲁德诺",
		TitleCode: "b_brodno",
	}
}
