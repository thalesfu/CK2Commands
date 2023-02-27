package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉扬SarayanBarony struct {
	feud.BaseBarony
}

var BSarayan萨拉扬 feud.Barony = &萨拉扬SarayanBarony{}

func init() {
    f := BSarayan萨拉扬.(*萨拉扬SarayanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarayan",
		TitleName: "萨拉扬",
		TitleCode: "b_sarayan",
	}
}
