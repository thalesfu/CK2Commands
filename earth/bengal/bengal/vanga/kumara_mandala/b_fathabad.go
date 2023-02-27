package kumara_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法特哈巴德FathabadBarony struct {
	feud.BaseBarony
}

var BFathabad法特哈巴德 feud.Barony = &法特哈巴德FathabadBarony{}

func init() {
    f := BFathabad法特哈巴德.(*法特哈巴德FathabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fathabad",
		TitleName: "法特哈巴德",
		TitleCode: "b_fathabad",
	}
}
