package iasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔勒乌HarlauBarony struct {
	feud.BaseBarony
}

var BHarlau赫尔勒乌 feud.Barony = &赫尔勒乌HarlauBarony{}

func init() {
    f := BHarlau赫尔勒乌.(*赫尔勒乌HarlauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harlau",
		TitleName: "赫尔勒乌",
		TitleCode: "b_harlau",
	}
}
