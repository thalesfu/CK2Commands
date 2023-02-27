package dwarasamudra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔西盖雷ArasikereBarony struct {
	feud.BaseBarony
}

var BArasikere阿尔西盖雷 feud.Barony = &阿尔西盖雷ArasikereBarony{}

func init() {
    f := BArasikere阿尔西盖雷.(*阿尔西盖雷ArasikereBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arasikere",
		TitleName: "阿尔西盖雷",
		TitleCode: "b_arasikere",
	}
}
