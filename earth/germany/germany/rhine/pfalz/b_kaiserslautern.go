package pfalz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯撒斯劳滕KaiserslauternBarony struct {
	feud.BaseBarony
}

var BKaiserslautern凯撒斯劳滕 feud.Barony = &凯撒斯劳滕KaiserslauternBarony{}

func init() {
	f := BKaiserslautern凯撒斯劳滕.(*凯撒斯劳滕KaiserslauternBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaiserslautern",
		TitleName: "凯撒斯劳滕",
		TitleCode: "b_kaiserslautern",
	}
}
