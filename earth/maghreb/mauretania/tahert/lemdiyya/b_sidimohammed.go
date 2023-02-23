package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪穆罕默德SidimohammedBarony struct {
	feud.BaseBarony
}

var BSidimohammed西迪穆罕默德 feud.Barony = &西迪穆罕默德SidimohammedBarony{}

func init() {
	f := BSidimohammed西迪穆罕默德.(*西迪穆罕默德SidimohammedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidimohammed",
		TitleName: "西迪穆罕默德",
		TitleCode: "b_sidimohammed",
	}
}
