package swetaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽吉KharkiBarony struct {
	feud.BaseBarony
}

var BKharki伽吉 feud.Barony = &伽吉KharkiBarony{}

func init() {
    f := BKharki伽吉.(*伽吉KharkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharki",
		TitleName: "伽吉",
		TitleCode: "b_kharki",
	}
}
