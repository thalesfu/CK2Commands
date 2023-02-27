package qalqut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘努尔KannurBarony struct {
	feud.BaseBarony
}

var BKannur甘努尔 feud.Barony = &甘努尔KannurBarony{}

func init() {
    f := BKannur甘努尔.(*甘努尔KannurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kannur",
		TitleName: "甘努尔",
		TitleCode: "b_kannur",
	}
}
