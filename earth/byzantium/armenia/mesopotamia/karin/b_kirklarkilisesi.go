package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基克拉克吉利瑟斯KirklarkilisesiBarony struct {
	feud.BaseBarony
}

var BKirklarkilisesi基克拉克吉利瑟斯 feud.Barony = &基克拉克吉利瑟斯KirklarkilisesiBarony{}

func init() {
	f := BKirklarkilisesi基克拉克吉利瑟斯.(*基克拉克吉利瑟斯KirklarkilisesiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirklarkilisesi",
		TitleName: "基克拉克吉利瑟斯",
		TitleCode: "b_kirklarkilisesi",
	}
}
