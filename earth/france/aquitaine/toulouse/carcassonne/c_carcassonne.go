package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CarcassonneCounty interface {
    feud.County
    BAlet阿莱() 	feud.Barony
    BCabaret卡巴雷() 	feud.Barony
    BCarcassonne卡尔卡松() 	feud.Barony
    BLagrasse拉格拉斯() 	feud.Barony
    BLastours拉斯图尔() 	feud.Barony
    BMinerve米内尔夫() 	feud.Barony
    BSaissac赛萨克() 	feud.Barony
    BTermes泰尔姆() 	feud.Barony
}

type 卡尔卡松CarcassonneCounty struct {
	feud.BaseCounty
	阿莱Alet 	feud.Barony
	卡巴雷Cabaret 	feud.Barony
	卡尔卡松Carcassonne 	feud.Barony
	拉格拉斯Lagrasse 	feud.Barony
	拉斯图尔Lastours 	feud.Barony
	米内尔夫Minerve 	feud.Barony
	赛萨克Saissac 	feud.Barony
	泰尔姆Termes 	feud.Barony
}

func (c *卡尔卡松CarcassonneCounty) BAlet阿莱() feud.Barony {
	return c.阿莱Alet
}
    
func (c *卡尔卡松CarcassonneCounty) BCabaret卡巴雷() feud.Barony {
	return c.卡巴雷Cabaret
}
    
func (c *卡尔卡松CarcassonneCounty) BCarcassonne卡尔卡松() feud.Barony {
	return c.卡尔卡松Carcassonne
}
    
func (c *卡尔卡松CarcassonneCounty) BLagrasse拉格拉斯() feud.Barony {
	return c.拉格拉斯Lagrasse
}
    
func (c *卡尔卡松CarcassonneCounty) BLastours拉斯图尔() feud.Barony {
	return c.拉斯图尔Lastours
}
    
func (c *卡尔卡松CarcassonneCounty) BMinerve米内尔夫() feud.Barony {
	return c.米内尔夫Minerve
}
    
func (c *卡尔卡松CarcassonneCounty) BSaissac赛萨克() feud.Barony {
	return c.赛萨克Saissac
}
    
func (c *卡尔卡松CarcassonneCounty) BTermes泰尔姆() feud.Barony {
	return c.泰尔姆Termes
}
    
var CCarcassonne卡尔卡松 CarcassonneCounty = &卡尔卡松CarcassonneCounty{}

func init() {
	f := CCarcassonne卡尔卡松.(*卡尔卡松CarcassonneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "213",
		Title:     "carcassonne",
		TitleName: "卡尔卡松",
		TitleCode: "c_carcassonne",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿莱Alet = BAlet阿莱
	f.阿莱Alet.SetParent(f)
	
	f.卡巴雷Cabaret = BCabaret卡巴雷
	f.卡巴雷Cabaret.SetParent(f)
	
	f.卡尔卡松Carcassonne = BCarcassonne卡尔卡松
	f.卡尔卡松Carcassonne.SetParent(f)
	
	f.拉格拉斯Lagrasse = BLagrasse拉格拉斯
	f.拉格拉斯Lagrasse.SetParent(f)
	
	f.拉斯图尔Lastours = BLastours拉斯图尔
	f.拉斯图尔Lastours.SetParent(f)
	
	f.米内尔夫Minerve = BMinerve米内尔夫
	f.米内尔夫Minerve.SetParent(f)
	
	f.赛萨克Saissac = BSaissac赛萨克
	f.赛萨克Saissac.SetParent(f)
	
	f.泰尔姆Termes = BTermes泰尔姆
	f.泰尔姆Termes.SetParent(f)
	
}
