package dadhipadra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀提波陀罗DadhipadraBarony struct {
	feud.BaseBarony
}

var BDadhipadra陀提波陀罗 feud.Barony = &陀提波陀罗DadhipadraBarony{}

func init() {
	f := BDadhipadra陀提波陀罗.(*陀提波陀罗DadhipadraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dadhipadra",
		TitleName: "陀提波陀罗",
		TitleCode: "b_dadhipadra",
	}
}
