package eastern_sayan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔然ArzhanBarony struct {
	feud.BaseBarony
}

var BArzhan阿尔然 feud.Barony = &阿尔然ArzhanBarony{}

func init() {
    f := BArzhan阿尔然.(*阿尔然ArzhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arzhan",
		TitleName: "阿尔然",
		TitleCode: "b_arzhan",
	}
}
