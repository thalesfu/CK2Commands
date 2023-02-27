package rusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁萨RusaBarony struct {
	feud.BaseBarony
}

var BRusa鲁萨 feud.Barony = &鲁萨RusaBarony{}

func init() {
    f := BRusa鲁萨.(*鲁萨RusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rusa",
		TitleName: "鲁萨",
		TitleCode: "b_rusa",
	}
}
