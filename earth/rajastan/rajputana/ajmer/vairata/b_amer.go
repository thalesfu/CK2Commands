package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿墨尔AmerBarony struct {
	feud.BaseBarony
}

var BAmer阿墨尔 feud.Barony = &阿墨尔AmerBarony{}

func init() {
	f := BAmer阿墨尔.(*阿墨尔AmerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amer",
		TitleName: "阿墨尔",
		TitleCode: "b_amer",
	}
}
