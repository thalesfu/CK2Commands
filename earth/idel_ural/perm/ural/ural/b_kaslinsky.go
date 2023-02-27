package ural

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯林斯基KaslinskyBarony struct {
	feud.BaseBarony
}

var BKaslinsky卡斯林斯基 feud.Barony = &卡斯林斯基KaslinskyBarony{}

func init() {
    f := BKaslinsky卡斯林斯基.(*卡斯林斯基KaslinskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaslinsky",
		TitleName: "卡斯林斯基",
		TitleCode: "b_kaslinsky",
	}
}
