package amous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿玛乌AmaousBarony struct {
	feud.BaseBarony
}

var BAmaous阿玛乌 feud.Barony = &阿玛乌AmaousBarony{}

func init() {
    f := BAmaous阿玛乌.(*阿玛乌AmaousBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amaous",
		TitleName: "阿玛乌",
		TitleCode: "b_amaous",
	}
}
