package iona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里纳古尔ArinagourBarony struct {
	feud.BaseBarony
}

var BArinagour阿里纳古尔 feud.Barony = &阿里纳古尔ArinagourBarony{}

func init() {
    f := BArinagour阿里纳古尔.(*阿里纳古尔ArinagourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arinagour",
		TitleName: "阿里纳古尔",
		TitleCode: "b_arinagour",
	}
}
