package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉哈尔AlajarBarony struct {
	feud.BaseBarony
}

var BAlajar阿拉哈尔 feud.Barony = &阿拉哈尔AlajarBarony{}

func init() {
	f := BAlajar阿拉哈尔.(*阿拉哈尔AlajarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alajar",
		TitleName: "阿拉哈尔",
		TitleCode: "b_alajar",
	}
}
