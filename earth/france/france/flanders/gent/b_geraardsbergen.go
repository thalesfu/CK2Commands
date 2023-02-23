package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉尔兹贝亨GeraardsbergenBarony struct {
	feud.BaseBarony
}

var BGeraardsbergen赫拉尔兹贝亨 feud.Barony = &赫拉尔兹贝亨GeraardsbergenBarony{}

func init() {
	f := BGeraardsbergen赫拉尔兹贝亨.(*赫拉尔兹贝亨GeraardsbergenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geraardsbergen",
		TitleName: "赫拉尔兹贝亨",
		TitleCode: "b_geraardsbergen",
	}
}
