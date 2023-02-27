package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特拉基KastrakiBarony struct {
	feud.BaseBarony
}

var BKastraki卡斯特拉基 feud.Barony = &卡斯特拉基KastrakiBarony{}

func init() {
    f := BKastraki卡斯特拉基.(*卡斯特拉基KastrakiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kastraki",
		TitleName: "卡斯特拉基",
		TitleCode: "b_kastraki",
	}
}
