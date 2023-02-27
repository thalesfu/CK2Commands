package chauragarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChauragarhCounty interface {
    feud.County
    BBarman婆罗曼() 	feud.Barony
    BBohani菩诃尼() 	feud.Barony
    BChauragarh招罗迦罗() 	feud.Barony
    BHainsasar含娑萨() 	feud.Barony
    BSirsali斯娑梨() 	feud.Barony
    BSiruvachur斯卢伐邱() 	feud.Barony
    BTingra鼎揭罗() 	feud.Barony
}

type 招罗迦罗ChauragarhCounty struct {
	feud.BaseCounty
	婆罗曼Barman 	feud.Barony
	菩诃尼Bohani 	feud.Barony
	招罗迦罗Chauragarh 	feud.Barony
	含娑萨Hainsasar 	feud.Barony
	斯娑梨Sirsali 	feud.Barony
	斯卢伐邱Siruvachur 	feud.Barony
	鼎揭罗Tingra 	feud.Barony
}

func (c *招罗迦罗ChauragarhCounty) BBarman婆罗曼() feud.Barony {
	return c.婆罗曼Barman
}
    
func (c *招罗迦罗ChauragarhCounty) BBohani菩诃尼() feud.Barony {
	return c.菩诃尼Bohani
}
    
func (c *招罗迦罗ChauragarhCounty) BChauragarh招罗迦罗() feud.Barony {
	return c.招罗迦罗Chauragarh
}
    
func (c *招罗迦罗ChauragarhCounty) BHainsasar含娑萨() feud.Barony {
	return c.含娑萨Hainsasar
}
    
func (c *招罗迦罗ChauragarhCounty) BSirsali斯娑梨() feud.Barony {
	return c.斯娑梨Sirsali
}
    
func (c *招罗迦罗ChauragarhCounty) BSiruvachur斯卢伐邱() feud.Barony {
	return c.斯卢伐邱Siruvachur
}
    
func (c *招罗迦罗ChauragarhCounty) BTingra鼎揭罗() feud.Barony {
	return c.鼎揭罗Tingra
}
    
var CChauragarh招罗迦罗 ChauragarhCounty = &招罗迦罗ChauragarhCounty{}

func init() {
	f := CChauragarh招罗迦罗.(*招罗迦罗ChauragarhCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1164",
		Title:     "chauragarh",
		TitleName: "招罗迦罗",
		TitleCode: "c_chauragarh",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆罗曼Barman = BBarman婆罗曼
	f.婆罗曼Barman.SetParent(f)
	
	f.菩诃尼Bohani = BBohani菩诃尼
	f.菩诃尼Bohani.SetParent(f)
	
	f.招罗迦罗Chauragarh = BChauragarh招罗迦罗
	f.招罗迦罗Chauragarh.SetParent(f)
	
	f.含娑萨Hainsasar = BHainsasar含娑萨
	f.含娑萨Hainsasar.SetParent(f)
	
	f.斯娑梨Sirsali = BSirsali斯娑梨
	f.斯娑梨Sirsali.SetParent(f)
	
	f.斯卢伐邱Siruvachur = BSiruvachur斯卢伐邱
	f.斯卢伐邱Siruvachur.SetParent(f)
	
	f.鼎揭罗Tingra = BTingra鼎揭罗
	f.鼎揭罗Tingra.SetParent(f)
	
}
