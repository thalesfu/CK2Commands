package ujjayini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃迦丽湿伐罗MahakaleshwarBarony struct {
	feud.BaseBarony
}

var BMahakaleshwar摩诃迦丽湿伐罗 feud.Barony = &摩诃迦丽湿伐罗MahakaleshwarBarony{}

func init() {
	f := BMahakaleshwar摩诃迦丽湿伐罗.(*摩诃迦丽湿伐罗MahakaleshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahakaleshwar",
		TitleName: "摩诃迦丽湿伐罗",
		TitleCode: "b_mahakaleshwar",
	}
}
