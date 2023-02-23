package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布罗伯里BroburyBarony struct {
	feud.BaseBarony
}

var BBrobury布罗伯里 feud.Barony = &布罗伯里BroburyBarony{}

func init() {
	f := BBrobury布罗伯里.(*布罗伯里BroburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brobury",
		TitleName: "布罗伯里",
		TitleCode: "b_brobury",
	}
}
