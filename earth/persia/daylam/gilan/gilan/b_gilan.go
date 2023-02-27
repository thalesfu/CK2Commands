package gilan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉兰GilanBarony struct {
	feud.BaseBarony
}

var BGilan吉兰 feud.Barony = &吉兰GilanBarony{}

func init() {
    f := BGilan吉兰.(*吉兰GilanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gilan",
		TitleName: "吉兰",
		TitleCode: "b_gilan",
	}
}
