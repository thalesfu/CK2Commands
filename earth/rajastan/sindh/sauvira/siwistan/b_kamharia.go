package siwistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘诃梨KamhariaBarony struct {
	feud.BaseBarony
}

var BKamharia甘诃梨 feud.Barony = &甘诃梨KamhariaBarony{}

func init() {
    f := BKamharia甘诃梨.(*甘诃梨KamhariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamharia",
		TitleName: "甘诃梨",
		TitleCode: "b_kamharia",
	}
}
