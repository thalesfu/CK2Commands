package vozviahel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别尔基切夫BerdytchivBarony struct {
	feud.BaseBarony
}

var BBerdytchiv别尔基切夫 feud.Barony = &别尔基切夫BerdytchivBarony{}

func init() {
    f := BBerdytchiv别尔基切夫.(*别尔基切夫BerdytchivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berdytchiv",
		TitleName: "别尔基切夫",
		TitleCode: "b_berdytchiv",
	}
}
