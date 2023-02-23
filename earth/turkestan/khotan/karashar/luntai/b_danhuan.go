package luntai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 单桓DanhuanBarony struct {
	feud.BaseBarony
}

var BDanhuan单桓 feud.Barony = &单桓DanhuanBarony{}

func init() {
	f := BDanhuan单桓.(*单桓DanhuanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danhuan",
		TitleName: "单桓",
		TitleCode: "b_danhuan",
	}
}
