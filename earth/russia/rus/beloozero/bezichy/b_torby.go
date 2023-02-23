package bezichy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔贝TorbyBarony struct {
	feud.BaseBarony
}

var BTorby托尔贝 feud.Barony = &托尔贝TorbyBarony{}

func init() {
	f := BTorby托尔贝.(*托尔贝TorbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torby",
		TitleName: "托尔贝",
		TitleCode: "b_torby",
	}
}
