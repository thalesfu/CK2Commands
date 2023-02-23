package soso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库瓦奥KoioBarony struct {
	feud.BaseBarony
}

var BKoio库瓦奥 feud.Barony = &库瓦奥KoioBarony{}

func init() {
	f := BKoio库瓦奥.(*库瓦奥KoioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koio",
		TitleName: "库瓦奥",
		TitleCode: "b_koio",
	}
}
