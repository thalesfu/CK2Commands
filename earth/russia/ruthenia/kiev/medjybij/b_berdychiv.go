package medjybij

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别尔基切夫BerdychivBarony struct {
	feud.BaseBarony
}

var BBerdychiv别尔基切夫 feud.Barony = &别尔基切夫BerdychivBarony{}

func init() {
    f := BBerdychiv别尔基切夫.(*别尔基切夫BerdychivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berdychiv",
		TitleName: "别尔基切夫",
		TitleCode: "b_berdychiv",
	}
}
