package tahoua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔明杜杜Kurmin_dodoBarony struct {
	feud.BaseBarony
}

var BKurmin_dodo库尔明杜杜 feud.Barony = &库尔明杜杜Kurmin_dodoBarony{}

func init() {
    f := BKurmin_dodo库尔明杜杜.(*库尔明杜杜Kurmin_dodoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurmin_dodo",
		TitleName: "库尔明杜杜",
		TitleCode: "b_kurmin_dodo",
	}
}
