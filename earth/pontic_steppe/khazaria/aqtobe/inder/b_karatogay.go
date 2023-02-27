package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉托盖KaratogayBarony struct {
	feud.BaseBarony
}

var BKaratogay卡拉托盖 feud.Barony = &卡拉托盖KaratogayBarony{}

func init() {
    f := BKaratogay卡拉托盖.(*卡拉托盖KaratogayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karatogay",
		TitleName: "卡拉托盖",
		TitleCode: "b_karatogay",
	}
}
