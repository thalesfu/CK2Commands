package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔克马尔AlkmaarBarony struct {
	feud.BaseBarony
}

var BAlkmaar阿尔克马尔 feud.Barony = &阿尔克马尔AlkmaarBarony{}

func init() {
	f := BAlkmaar阿尔克马尔.(*阿尔克马尔AlkmaarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alkmaar",
		TitleName: "阿尔克马尔",
		TitleCode: "b_alkmaar",
	}
}
