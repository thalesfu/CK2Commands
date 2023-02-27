package mtskheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝布里斯齐赫BebristsikheBarony struct {
	feud.BaseBarony
}

var BBebristsikhe贝布里斯齐赫 feud.Barony = &贝布里斯齐赫BebristsikheBarony{}

func init() {
    f := BBebristsikhe贝布里斯齐赫.(*贝布里斯齐赫BebristsikheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bebristsikhe",
		TitleName: "贝布里斯齐赫",
		TitleCode: "b_bebristsikhe",
	}
}
