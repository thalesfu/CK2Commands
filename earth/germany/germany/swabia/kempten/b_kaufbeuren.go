package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考夫博伊伦KaufbeurenBarony struct {
	feud.BaseBarony
}

var BKaufbeuren考夫博伊伦 feud.Barony = &考夫博伊伦KaufbeurenBarony{}

func init() {
	f := BKaufbeuren考夫博伊伦.(*考夫博伊伦KaufbeurenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaufbeuren",
		TitleName: "考夫博伊伦",
		TitleCode: "b_kaufbeuren",
	}
}
