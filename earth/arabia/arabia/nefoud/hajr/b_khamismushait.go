package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海米斯穆谢特KhamismushaitBarony struct {
	feud.BaseBarony
}

var BKhamismushait海米斯穆谢特 feud.Barony = &海米斯穆谢特KhamismushaitBarony{}

func init() {
	f := BKhamismushait海米斯穆谢特.(*海米斯穆谢特KhamismushaitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khamismushait",
		TitleName: "海米斯穆谢特",
		TitleCode: "b_khamismushait",
	}
}
