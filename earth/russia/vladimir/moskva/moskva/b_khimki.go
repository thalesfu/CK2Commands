package moskva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希姆基KhimkiBarony struct {
	feud.BaseBarony
}

var BKhimki希姆基 feud.Barony = &希姆基KhimkiBarony{}

func init() {
	f := BKhimki希姆基.(*希姆基KhimkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khimki",
		TitleName: "希姆基",
		TitleCode: "b_khimki",
	}
}
