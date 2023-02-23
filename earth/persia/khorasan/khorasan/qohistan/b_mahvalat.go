package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马瓦拉特MahvalatBarony struct {
	feud.BaseBarony
}

var BMahvalat马瓦拉特 feud.Barony = &马瓦拉特MahvalatBarony{}

func init() {
	f := BMahvalat马瓦拉特.(*马瓦拉特MahvalatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahvalat",
		TitleName: "马瓦拉特",
		TitleCode: "b_mahvalat",
	}
}
