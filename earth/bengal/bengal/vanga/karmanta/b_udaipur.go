package karmanta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优陀耶补罗UdaipurBarony struct {
	feud.BaseBarony
}

var BUdaipur优陀耶补罗 feud.Barony = &优陀耶补罗UdaipurBarony{}

func init() {
	f := BUdaipur优陀耶补罗.(*优陀耶补罗UdaipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udaipur",
		TitleName: "优陀耶补罗",
		TitleCode: "b_udaipur",
	}
}
