package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣布里吉德St_brigitBarony struct {
	feud.BaseBarony
}

var BSt_brigit圣布里吉德 feud.Barony = &圣布里吉德St_brigitBarony{}

func init() {
    f := BSt_brigit圣布里吉德.(*圣布里吉德St_brigitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_brigit",
		TitleName: "圣布里吉德",
		TitleCode: "b_st_brigit",
	}
}
