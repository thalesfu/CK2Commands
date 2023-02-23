package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀瑟尔QoserBarony struct {
	feud.BaseBarony
}

var BQoser喀瑟尔 feud.Barony = &喀瑟尔QoserBarony{}

func init() {
	f := BQoser喀瑟尔.(*喀瑟尔QoserBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qoser",
		TitleName: "喀瑟尔",
		TitleCode: "b_qoser",
	}
}
