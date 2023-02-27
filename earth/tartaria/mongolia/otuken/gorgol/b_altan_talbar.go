package gorgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿勒坦塔勒巴尔Altan_talbarBarony struct {
	feud.BaseBarony
}

var BAltan_talbar阿勒坦塔勒巴尔 feud.Barony = &阿勒坦塔勒巴尔Altan_talbarBarony{}

func init() {
    f := BAltan_talbar阿勒坦塔勒巴尔.(*阿勒坦塔勒巴尔Altan_talbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altan_talbar",
		TitleName: "阿勒坦塔勒巴尔",
		TitleCode: "b_altan_talbar",
	}
}
