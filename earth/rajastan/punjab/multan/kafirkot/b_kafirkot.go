package kafirkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡菲尔科特KafirkotBarony struct {
	feud.BaseBarony
}

var BKafirkot卡菲尔科特 feud.Barony = &卡菲尔科特KafirkotBarony{}

func init() {
	f := BKafirkot卡菲尔科特.(*卡菲尔科特KafirkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kafirkot",
		TitleName: "卡菲尔科特",
		TitleCode: "b_kafirkot",
	}
}
