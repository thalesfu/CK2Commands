package smaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔维斯塔AlvestaBarony struct {
	feud.BaseBarony
}

var BAlvesta阿尔维斯塔 feud.Barony = &阿尔维斯塔AlvestaBarony{}

func init() {
    f := BAlvesta阿尔维斯塔.(*阿尔维斯塔AlvestaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alvesta",
		TitleName: "阿尔维斯塔",
		TitleCode: "b_alvesta",
	}
}
