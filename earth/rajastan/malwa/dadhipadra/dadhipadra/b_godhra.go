package dadhipadra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈持罗GodhraBarony struct {
	feud.BaseBarony
}

var BGodhra戈持罗 feud.Barony = &戈持罗GodhraBarony{}

func init() {
	f := BGodhra戈持罗.(*戈持罗GodhraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "godhra",
		TitleName: "戈持罗",
		TitleCode: "b_godhra",
	}
}
