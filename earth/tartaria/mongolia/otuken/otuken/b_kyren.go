package otuken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基伦KyrenBarony struct {
	feud.BaseBarony
}

var BKyren基伦 feud.Barony = &基伦KyrenBarony{}

func init() {
	f := BKyren基伦.(*基伦KyrenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyren",
		TitleName: "基伦",
		TitleCode: "b_kyren",
	}
}
