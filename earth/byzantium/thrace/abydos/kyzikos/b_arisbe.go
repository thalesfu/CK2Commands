package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里斯贝ArisbeBarony struct {
	feud.BaseBarony
}

var BArisbe阿里斯贝 feud.Barony = &阿里斯贝ArisbeBarony{}

func init() {
	f := BArisbe阿里斯贝.(*阿里斯贝ArisbeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arisbe",
		TitleName: "阿里斯贝",
		TitleCode: "b_arisbe",
	}
}
