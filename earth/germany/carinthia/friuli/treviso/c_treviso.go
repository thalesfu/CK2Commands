package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrevisoCounty interface {
    feud.County
    BAsola阿索拉() 	feud.Barony
    BBassano巴萨诺() 	feud.Barony
    BCastelfranco卡斯泰尔夫兰科() 	feud.Barony
    BCeneda卡内达() 	feud.Barony
    BCitadella奇塔代拉() 	feud.Barony
    BOderzo奥代尔佐() 	feud.Barony
    BPaese帕埃塞() 	feud.Barony
    BTreviso特雷维索() 	feud.Barony
}

type 特雷维索TrevisoCounty struct {
	feud.BaseCounty
	阿索拉Asola 	feud.Barony
	巴萨诺Bassano 	feud.Barony
	卡斯泰尔夫兰科Castelfranco 	feud.Barony
	卡内达Ceneda 	feud.Barony
	奇塔代拉Citadella 	feud.Barony
	奥代尔佐Oderzo 	feud.Barony
	帕埃塞Paese 	feud.Barony
	特雷维索Treviso 	feud.Barony
}

func (c *特雷维索TrevisoCounty) BAsola阿索拉() feud.Barony {
	return c.阿索拉Asola
}
    
func (c *特雷维索TrevisoCounty) BBassano巴萨诺() feud.Barony {
	return c.巴萨诺Bassano
}
    
func (c *特雷维索TrevisoCounty) BCastelfranco卡斯泰尔夫兰科() feud.Barony {
	return c.卡斯泰尔夫兰科Castelfranco
}
    
func (c *特雷维索TrevisoCounty) BCeneda卡内达() feud.Barony {
	return c.卡内达Ceneda
}
    
func (c *特雷维索TrevisoCounty) BCitadella奇塔代拉() feud.Barony {
	return c.奇塔代拉Citadella
}
    
func (c *特雷维索TrevisoCounty) BOderzo奥代尔佐() feud.Barony {
	return c.奥代尔佐Oderzo
}
    
func (c *特雷维索TrevisoCounty) BPaese帕埃塞() feud.Barony {
	return c.帕埃塞Paese
}
    
func (c *特雷维索TrevisoCounty) BTreviso特雷维索() feud.Barony {
	return c.特雷维索Treviso
}
    
var CTreviso特雷维索 TrevisoCounty = &特雷维索TrevisoCounty{}

func init() {
	f := CTreviso特雷维索.(*特雷维索TrevisoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "357",
		Title:     "treviso",
		TitleName: "特雷维索",
		TitleCode: "c_treviso",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿索拉Asola = BAsola阿索拉
	f.阿索拉Asola.SetParent(f)
	
	f.巴萨诺Bassano = BBassano巴萨诺
	f.巴萨诺Bassano.SetParent(f)
	
	f.卡斯泰尔夫兰科Castelfranco = BCastelfranco卡斯泰尔夫兰科
	f.卡斯泰尔夫兰科Castelfranco.SetParent(f)
	
	f.卡内达Ceneda = BCeneda卡内达
	f.卡内达Ceneda.SetParent(f)
	
	f.奇塔代拉Citadella = BCitadella奇塔代拉
	f.奇塔代拉Citadella.SetParent(f)
	
	f.奥代尔佐Oderzo = BOderzo奥代尔佐
	f.奥代尔佐Oderzo.SetParent(f)
	
	f.帕埃塞Paese = BPaese帕埃塞
	f.帕埃塞Paese.SetParent(f)
	
	f.特雷维索Treviso = BTreviso特雷维索
	f.特雷维索Treviso.SetParent(f)
	
}
