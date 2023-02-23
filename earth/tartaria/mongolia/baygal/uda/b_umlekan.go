package uda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆列坎UmlekanBarony struct {
	feud.BaseBarony
}

var BUmlekan乌姆列坎 feud.Barony = &乌姆列坎UmlekanBarony{}

func init() {
	f := BUmlekan乌姆列坎.(*乌姆列坎UmlekanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umlekan",
		TitleName: "乌姆列坎",
		TitleCode: "b_umlekan",
	}
}
