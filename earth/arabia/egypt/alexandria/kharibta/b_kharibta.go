package kharibta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海里卜塔KharibtaBarony struct {
	feud.BaseBarony
}

var BKharibta海里卜塔 feud.Barony = &海里卜塔KharibtaBarony{}

func init() {
	f := BKharibta海里卜塔.(*海里卜塔KharibtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharibta",
		TitleName: "海里卜塔",
		TitleCode: "b_kharibta",
	}
}
