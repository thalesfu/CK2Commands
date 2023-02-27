package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈卡贝德HaykaberdBarony struct {
	feud.BaseBarony
}

var BHaykaberd哈卡贝德 feud.Barony = &哈卡贝德HaykaberdBarony{}

func init() {
    f := BHaykaberd哈卡贝德.(*哈卡贝德HaykaberdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haykaberd",
		TitleName: "哈卡贝德",
		TitleCode: "b_haykaberd",
	}
}
