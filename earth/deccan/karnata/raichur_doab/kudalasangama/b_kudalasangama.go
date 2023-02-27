package kudalasangama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俱荼罗僧伽摩KudalasangamaBarony struct {
	feud.BaseBarony
}

var BKudalasangama俱荼罗僧伽摩 feud.Barony = &俱荼罗僧伽摩KudalasangamaBarony{}

func init() {
    f := BKudalasangama俱荼罗僧伽摩.(*俱荼罗僧伽摩KudalasangamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kudalasangama",
		TitleName: "俱荼罗僧伽摩",
		TitleCode: "b_kudalasangama",
	}
}
