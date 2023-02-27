package karghalik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丁坑口DingkengkouBarony struct {
	feud.BaseBarony
}

var BDingkengkou丁坑口 feud.Barony = &丁坑口DingkengkouBarony{}

func init() {
    f := BDingkengkou丁坑口.(*丁坑口DingkengkouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dingkengkou",
		TitleName: "丁坑口",
		TitleCode: "b_dingkengkou",
	}
}
