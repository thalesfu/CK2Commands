package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PundravardhanaCounty interface {
	feud.County
	BGhoraghat瞿罗伽吒() feud.Barony
	BMahasthangarh摩诃萨傥那姞利呬() feud.Barony
	BPundravardhana奔那伐弹那() feud.Barony
}

type 奔那伐弹那PundravardhanaCounty struct {
	feud.BaseCounty
	瞿罗伽吒Ghoraghat         feud.Barony
	摩诃萨傥那姞利呬Mahasthangarh feud.Barony
	奔那伐弹那Pundravardhana   feud.Barony
}

func (c *奔那伐弹那PundravardhanaCounty) BGhoraghat瞿罗伽吒() feud.Barony {
	return c.瞿罗伽吒Ghoraghat
}

func (c *奔那伐弹那PundravardhanaCounty) BMahasthangarh摩诃萨傥那姞利呬() feud.Barony {
	return c.摩诃萨傥那姞利呬Mahasthangarh
}

func (c *奔那伐弹那PundravardhanaCounty) BPundravardhana奔那伐弹那() feud.Barony {
	return c.奔那伐弹那Pundravardhana
}

var CPundravardhana奔那伐弹那 PundravardhanaCounty = &奔那伐弹那PundravardhanaCounty{}

func init() {
	f := CPundravardhana奔那伐弹那.(*奔那伐弹那PundravardhanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1381",
		Title:     "pundravardhana",
		TitleName: "奔那伐弹那",
		TitleCode: "c_pundravardhana",
		Baronies:  map[string]feud.Barony{},
	}

	f.瞿罗伽吒Ghoraghat = BGhoraghat瞿罗伽吒
	f.瞿罗伽吒Ghoraghat.SetParent(f)

	f.摩诃萨傥那姞利呬Mahasthangarh = BMahasthangarh摩诃萨傥那姞利呬
	f.摩诃萨傥那姞利呬Mahasthangarh.SetParent(f)

	f.奔那伐弹那Pundravardhana = BPundravardhana奔那伐弹那
	f.奔那伐弹那Pundravardhana.SetParent(f)

}
