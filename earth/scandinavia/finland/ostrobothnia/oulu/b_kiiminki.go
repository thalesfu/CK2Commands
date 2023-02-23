package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基明基KiiminkiBarony struct {
	feud.BaseBarony
}

var BKiiminki基明基 feud.Barony = &基明基KiiminkiBarony{}

func init() {
	f := BKiiminki基明基.(*基明基KiiminkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiiminki",
		TitleName: "基明基",
		TitleCode: "b_kiiminki",
	}
}
