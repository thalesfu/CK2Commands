package udmurts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌德穆尔特UdmurtsBarony struct {
	feud.BaseBarony
}

var BUdmurts乌德穆尔特 feud.Barony = &乌德穆尔特UdmurtsBarony{}

func init() {
    f := BUdmurts乌德穆尔特.(*乌德穆尔特UdmurtsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udmurts",
		TitleName: "乌德穆尔特",
		TitleCode: "b_udmurts",
	}
}
