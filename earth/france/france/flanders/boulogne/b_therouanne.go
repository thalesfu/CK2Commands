package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰鲁阿讷TherouanneBarony struct {
	feud.BaseBarony
}

var BTherouanne泰鲁阿讷 feud.Barony = &泰鲁阿讷TherouanneBarony{}

func init() {
    f := BTherouanne泰鲁阿讷.(*泰鲁阿讷TherouanneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "therouanne",
		TitleName: "泰鲁阿讷",
		TitleCode: "b_therouanne",
	}
}
