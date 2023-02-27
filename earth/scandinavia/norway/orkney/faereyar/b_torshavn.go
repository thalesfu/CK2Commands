package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔斯豪恩TorshavnBarony struct {
	feud.BaseBarony
}

var BTorshavn托尔斯豪恩 feud.Barony = &托尔斯豪恩TorshavnBarony{}

func init() {
    f := BTorshavn托尔斯豪恩.(*托尔斯豪恩TorshavnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torshavn",
		TitleName: "托尔斯豪恩",
		TitleCode: "b_torshavn",
	}
}
