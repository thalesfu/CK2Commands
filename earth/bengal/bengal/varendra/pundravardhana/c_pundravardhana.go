package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PundravardhanaCounty interface {
    feud.County
    BBhimer_jangal毗咩罗旬迦罗() 	feud.Barony
    BGhoraghat瞿罗伽吒() 	feud.Barony
    BKhulnar_dhap丘罗土丘() 	feud.Barony
    BMahasthangarh摩诃萨傥那姞利呬() 	feud.Barony
    BPundravardhana奔那伐弹那() 	feud.Barony
    BSomapura_mahavihara苏摩补罗摩诃毗诃罗() 	feud.Barony
    BTotaram_panditer_dhap豆多蓝般提蒂罗土丘() 	feud.Barony
}

type 奔那伐弹那PundravardhanaCounty struct {
	feud.BaseCounty
	毗咩罗旬迦罗Bhimer_jangal 	feud.Barony
	瞿罗伽吒Ghoraghat 	feud.Barony
	丘罗土丘Khulnar_dhap 	feud.Barony
	摩诃萨傥那姞利呬Mahasthangarh 	feud.Barony
	奔那伐弹那Pundravardhana 	feud.Barony
	苏摩补罗摩诃毗诃罗Somapura_mahavihara 	feud.Barony
	豆多蓝般提蒂罗土丘Totaram_panditer_dhap 	feud.Barony
}

func (c *奔那伐弹那PundravardhanaCounty) BBhimer_jangal毗咩罗旬迦罗() feud.Barony {
	return c.毗咩罗旬迦罗Bhimer_jangal
}
    
func (c *奔那伐弹那PundravardhanaCounty) BGhoraghat瞿罗伽吒() feud.Barony {
	return c.瞿罗伽吒Ghoraghat
}
    
func (c *奔那伐弹那PundravardhanaCounty) BKhulnar_dhap丘罗土丘() feud.Barony {
	return c.丘罗土丘Khulnar_dhap
}
    
func (c *奔那伐弹那PundravardhanaCounty) BMahasthangarh摩诃萨傥那姞利呬() feud.Barony {
	return c.摩诃萨傥那姞利呬Mahasthangarh
}
    
func (c *奔那伐弹那PundravardhanaCounty) BPundravardhana奔那伐弹那() feud.Barony {
	return c.奔那伐弹那Pundravardhana
}
    
func (c *奔那伐弹那PundravardhanaCounty) BSomapura_mahavihara苏摩补罗摩诃毗诃罗() feud.Barony {
	return c.苏摩补罗摩诃毗诃罗Somapura_mahavihara
}
    
func (c *奔那伐弹那PundravardhanaCounty) BTotaram_panditer_dhap豆多蓝般提蒂罗土丘() feud.Barony {
	return c.豆多蓝般提蒂罗土丘Totaram_panditer_dhap
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

	f.毗咩罗旬迦罗Bhimer_jangal = BBhimer_jangal毗咩罗旬迦罗
	f.毗咩罗旬迦罗Bhimer_jangal.SetParent(f)
	
	f.瞿罗伽吒Ghoraghat = BGhoraghat瞿罗伽吒
	f.瞿罗伽吒Ghoraghat.SetParent(f)
	
	f.丘罗土丘Khulnar_dhap = BKhulnar_dhap丘罗土丘
	f.丘罗土丘Khulnar_dhap.SetParent(f)
	
	f.摩诃萨傥那姞利呬Mahasthangarh = BMahasthangarh摩诃萨傥那姞利呬
	f.摩诃萨傥那姞利呬Mahasthangarh.SetParent(f)
	
	f.奔那伐弹那Pundravardhana = BPundravardhana奔那伐弹那
	f.奔那伐弹那Pundravardhana.SetParent(f)
	
	f.苏摩补罗摩诃毗诃罗Somapura_mahavihara = BSomapura_mahavihara苏摩补罗摩诃毗诃罗
	f.苏摩补罗摩诃毗诃罗Somapura_mahavihara.SetParent(f)
	
	f.豆多蓝般提蒂罗土丘Totaram_panditer_dhap = BTotaram_panditer_dhap豆多蓝般提蒂罗土丘
	f.豆多蓝般提蒂罗土丘Totaram_panditer_dhap.SetParent(f)
	
}
