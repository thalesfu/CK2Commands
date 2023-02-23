package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔维斯尧尔ArvidsjaurBarony struct {
	feud.BaseBarony
}

var BArvidsjaur阿尔维斯尧尔 feud.Barony = &阿尔维斯尧尔ArvidsjaurBarony{}

func init() {
	f := BArvidsjaur阿尔维斯尧尔.(*阿尔维斯尧尔ArvidsjaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arvidsjaur",
		TitleName: "阿尔维斯尧尔",
		TitleCode: "b_arvidsjaur",
	}
}
