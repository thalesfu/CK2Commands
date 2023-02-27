package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫罗萨利亚MorosagliaBarony struct {
	feud.BaseBarony
}

var BMorosaglia莫罗萨利亚 feud.Barony = &莫罗萨利亚MorosagliaBarony{}

func init() {
    f := BMorosaglia莫罗萨利亚.(*莫罗萨利亚MorosagliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morosaglia",
		TitleName: "莫罗萨利亚",
		TitleCode: "b_morosaglia",
	}
}
