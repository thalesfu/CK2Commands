package guzgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯卜曼UsbumanBarony struct {
	feud.BaseBarony
}

var BUsbuman乌斯卜曼 feud.Barony = &乌斯卜曼UsbumanBarony{}

func init() {
	f := BUsbuman乌斯卜曼.(*乌斯卜曼UsbumanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usbuman",
		TitleName: "乌斯卜曼",
		TitleCode: "b_usbuman",
	}
}
