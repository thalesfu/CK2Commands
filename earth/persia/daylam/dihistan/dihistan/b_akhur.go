package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿胡尔AkhurBarony struct {
	feud.BaseBarony
}

var BAkhur阿胡尔 feud.Barony = &阿胡尔AkhurBarony{}

func init() {
	f := BAkhur阿胡尔.(*阿胡尔AkhurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhur",
		TitleName: "阿胡尔",
		TitleCode: "b_akhur",
	}
}
