package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔莫纳斯特拉雷亚尔AlmonasterlarealBarony struct {
	feud.BaseBarony
}

var BAlmonasterlareal阿尔莫纳斯特拉雷亚尔 feud.Barony = &阿尔莫纳斯特拉雷亚尔AlmonasterlarealBarony{}

func init() {
	f := BAlmonasterlareal阿尔莫纳斯特拉雷亚尔.(*阿尔莫纳斯特拉雷亚尔AlmonasterlarealBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almonasterlareal",
		TitleName: "阿尔莫纳斯特拉雷亚尔",
		TitleCode: "b_almonasterlareal",
	}
}
