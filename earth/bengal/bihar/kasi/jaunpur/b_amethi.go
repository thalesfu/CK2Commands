package jaunpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿咩底AmethiBarony struct {
	feud.BaseBarony
}

var BAmethi阿咩底 feud.Barony = &阿咩底AmethiBarony{}

func init() {
    f := BAmethi阿咩底.(*阿咩底AmethiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amethi",
		TitleName: "阿咩底",
		TitleCode: "b_amethi",
	}
}
