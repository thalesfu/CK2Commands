package saumur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜埃拉丰坦Doue_la_fontaineBarony struct {
	feud.BaseBarony
}

var BDoue_la_fontaine杜埃拉丰坦 feud.Barony = &杜埃拉丰坦Doue_la_fontaineBarony{}

func init() {
    f := BDoue_la_fontaine杜埃拉丰坦.(*杜埃拉丰坦Doue_la_fontaineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doue_la_fontaine",
		TitleName: "杜埃拉丰坦",
		TitleCode: "b_doue_la_fontaine",
	}
}
