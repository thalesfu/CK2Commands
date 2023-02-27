package kanchipuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿楞补罗毗AlampuraviBarony struct {
	feud.BaseBarony
}

var BAlampuravi阿楞补罗毗 feud.Barony = &阿楞补罗毗AlampuraviBarony{}

func init() {
    f := BAlampuravi阿楞补罗毗.(*阿楞补罗毗AlampuraviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alampuravi",
		TitleName: "阿楞补罗毗",
		TitleCode: "b_alampuravi",
	}
}
