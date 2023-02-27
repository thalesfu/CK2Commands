package laksmanavati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盘卡克PankakBarony struct {
	feud.BaseBarony
}

var BPankak盘卡克 feud.Barony = &盘卡克PankakBarony{}

func init() {
    f := BPankak盘卡克.(*盘卡克PankakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pankak",
		TitleName: "盘卡克",
		TitleCode: "b_pankak",
	}
}
