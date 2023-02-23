package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉特德拉特KalaatderatBarony struct {
	feud.BaseBarony
}

var BKalaatderat卡拉特德拉特 feud.Barony = &卡拉特德拉特KalaatderatBarony{}

func init() {
	f := BKalaatderat卡拉特德拉特.(*卡拉特德拉特KalaatderatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalaatderat",
		TitleName: "卡拉特德拉特",
		TitleCode: "b_kalaatderat",
	}
}
