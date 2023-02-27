package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TheodosiopolisCounty interface {
    feud.County
    BArgyropolis安古洛波利斯() 	feud.Barony
    BAskale阿什卡莱() 	feud.Barony
    BCitharizum基萨里宗() 	feud.Barony
    BOukhiti夸西提() 	feud.Barony
    BSatala萨塔拉() 	feud.Barony
    BTheodosiopolis狄奥多西波利斯() 	feud.Barony
    BThera特雷() 	feud.Barony
    BTortum托尔图姆() 	feud.Barony
}

type 狄奥多西波利斯TheodosiopolisCounty struct {
	feud.BaseCounty
	安古洛波利斯Argyropolis 	feud.Barony
	阿什卡莱Askale 	feud.Barony
	基萨里宗Citharizum 	feud.Barony
	夸西提Oukhiti 	feud.Barony
	萨塔拉Satala 	feud.Barony
	狄奥多西波利斯Theodosiopolis 	feud.Barony
	特雷Thera 	feud.Barony
	托尔图姆Tortum 	feud.Barony
}

func (c *狄奥多西波利斯TheodosiopolisCounty) BArgyropolis安古洛波利斯() feud.Barony {
	return c.安古洛波利斯Argyropolis
}
    
func (c *狄奥多西波利斯TheodosiopolisCounty) BAskale阿什卡莱() feud.Barony {
	return c.阿什卡莱Askale
}
    
func (c *狄奥多西波利斯TheodosiopolisCounty) BCitharizum基萨里宗() feud.Barony {
	return c.基萨里宗Citharizum
}
    
func (c *狄奥多西波利斯TheodosiopolisCounty) BOukhiti夸西提() feud.Barony {
	return c.夸西提Oukhiti
}
    
func (c *狄奥多西波利斯TheodosiopolisCounty) BSatala萨塔拉() feud.Barony {
	return c.萨塔拉Satala
}
    
func (c *狄奥多西波利斯TheodosiopolisCounty) BTheodosiopolis狄奥多西波利斯() feud.Barony {
	return c.狄奥多西波利斯Theodosiopolis
}
    
func (c *狄奥多西波利斯TheodosiopolisCounty) BThera特雷() feud.Barony {
	return c.特雷Thera
}
    
func (c *狄奥多西波利斯TheodosiopolisCounty) BTortum托尔图姆() feud.Barony {
	return c.托尔图姆Tortum
}
    
var CTheodosiopolis狄奥多西波利斯 TheodosiopolisCounty = &狄奥多西波利斯TheodosiopolisCounty{}

func init() {
	f := CTheodosiopolis狄奥多西波利斯.(*狄奥多西波利斯TheodosiopolisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "704",
		Title:     "theodosiopolis",
		TitleName: "狄奥多西波利斯",
		TitleCode: "c_theodosiopolis",
		Baronies:  map[string]feud.Barony{},
	}

	f.安古洛波利斯Argyropolis = BArgyropolis安古洛波利斯
	f.安古洛波利斯Argyropolis.SetParent(f)
	
	f.阿什卡莱Askale = BAskale阿什卡莱
	f.阿什卡莱Askale.SetParent(f)
	
	f.基萨里宗Citharizum = BCitharizum基萨里宗
	f.基萨里宗Citharizum.SetParent(f)
	
	f.夸西提Oukhiti = BOukhiti夸西提
	f.夸西提Oukhiti.SetParent(f)
	
	f.萨塔拉Satala = BSatala萨塔拉
	f.萨塔拉Satala.SetParent(f)
	
	f.狄奥多西波利斯Theodosiopolis = BTheodosiopolis狄奥多西波利斯
	f.狄奥多西波利斯Theodosiopolis.SetParent(f)
	
	f.特雷Thera = BThera特雷
	f.特雷Thera.SetParent(f)
	
	f.托尔图姆Tortum = BTortum托尔图姆
	f.托尔图姆Tortum.SetParent(f)
	
}
