package sripatha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀婆丽湿伐罗摩诃提婆Dholeshwar_mahadevBarony struct {
	feud.BaseBarony
}

var BDholeshwar_mahadev陀婆丽湿伐罗摩诃提婆 feud.Barony = &陀婆丽湿伐罗摩诃提婆Dholeshwar_mahadevBarony{}

func init() {
    f := BDholeshwar_mahadev陀婆丽湿伐罗摩诃提婆.(*陀婆丽湿伐罗摩诃提婆Dholeshwar_mahadevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dholeshwar_mahadev",
		TitleName: "陀婆丽湿伐罗摩诃提婆",
		TitleCode: "b_dholeshwar_mahadev",
	}
}
