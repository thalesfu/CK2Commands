package sagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 薄伽补罗BagpuraBarony struct {
	feud.BaseBarony
}

var BBagpura薄伽补罗 feud.Barony = &薄伽补罗BagpuraBarony{}

func init() {
	f := BBagpura薄伽补罗.(*薄伽补罗BagpuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagpura",
		TitleName: "薄伽补罗",
		TitleCode: "b_bagpura",
	}
}
