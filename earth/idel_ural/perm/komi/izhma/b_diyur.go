package izhma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季尤尔DiyurBarony struct {
	feud.BaseBarony
}

var BDiyur季尤尔 feud.Barony = &季尤尔DiyurBarony{}

func init() {
    f := BDiyur季尤尔.(*季尤尔DiyurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diyur",
		TitleName: "季尤尔",
		TitleCode: "b_diyur",
	}
}
