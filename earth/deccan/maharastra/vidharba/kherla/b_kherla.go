package kherla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 稽剌罗KherlaBarony struct {
	feud.BaseBarony
}

var BKherla稽剌罗 feud.Barony = &稽剌罗KherlaBarony{}

func init() {
	f := BKherla稽剌罗.(*稽剌罗KherlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kherla",
		TitleName: "稽剌罗",
		TitleCode: "b_kherla",
	}
}
