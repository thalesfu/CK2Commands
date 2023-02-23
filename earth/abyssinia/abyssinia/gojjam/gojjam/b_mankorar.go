package gojjam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼科拉尔MankorarBarony struct {
	feud.BaseBarony
}

var BMankorar曼科拉尔 feud.Barony = &曼科拉尔MankorarBarony{}

func init() {
	f := BMankorar曼科拉尔.(*曼科拉尔MankorarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mankorar",
		TitleName: "曼科拉尔",
		TitleCode: "b_mankorar",
	}
}
