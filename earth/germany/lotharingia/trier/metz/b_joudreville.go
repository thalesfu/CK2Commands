package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茹德勒维尔JoudrevilleBarony struct {
	feud.BaseBarony
}

var BJoudreville茹德勒维尔 feud.Barony = &茹德勒维尔JoudrevilleBarony{}

func init() {
	f := BJoudreville茹德勒维尔.(*茹德勒维尔JoudrevilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "joudreville",
		TitleName: "茹德勒维尔",
		TitleCode: "b_joudreville",
	}
}
