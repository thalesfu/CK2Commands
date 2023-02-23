package saptagrama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃那陀MahanadBarony struct {
	feud.BaseBarony
}

var BMahanad摩诃那陀 feud.Barony = &摩诃那陀MahanadBarony{}

func init() {
	f := BMahanad摩诃那陀.(*摩诃那陀MahanadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahanad",
		TitleName: "摩诃那陀",
		TitleCode: "b_mahanad",
	}
}
