package nellore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 郡稚伽补梨GundigapuriBarony struct {
	feud.BaseBarony
}

var BGundigapuri郡稚伽补梨 feud.Barony = &郡稚伽补梨GundigapuriBarony{}

func init() {
	f := BGundigapuri郡稚伽补梨.(*郡稚伽补梨GundigapuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gundigapuri",
		TitleName: "郡稚伽补梨",
		TitleCode: "b_gundigapuri",
	}
}
