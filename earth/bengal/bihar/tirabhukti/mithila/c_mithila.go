package mithila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MithilaCounty interface {
    feud.County
    BDarbhanga达罗傍伽() 	feud.Barony
    BHajipur诃耆城() 	feud.Barony
    BMithila弥梯罗() 	feud.Barony
    BSaligrama舍利伽罗摩() 	feud.Barony
    BSiwan锡万() 	feud.Barony
    BSugauna须乔那() 	feud.Barony
    BYavamajjhaka耶婆摩阇迦() 	feud.Barony
}

type 弥梯罗MithilaCounty struct {
	feud.BaseCounty
	达罗傍伽Darbhanga 	feud.Barony
	诃耆城Hajipur 	feud.Barony
	弥梯罗Mithila 	feud.Barony
	舍利伽罗摩Saligrama 	feud.Barony
	锡万Siwan 	feud.Barony
	须乔那Sugauna 	feud.Barony
	耶婆摩阇迦Yavamajjhaka 	feud.Barony
}

func (c *弥梯罗MithilaCounty) BDarbhanga达罗傍伽() feud.Barony {
	return c.达罗傍伽Darbhanga
}
    
func (c *弥梯罗MithilaCounty) BHajipur诃耆城() feud.Barony {
	return c.诃耆城Hajipur
}
    
func (c *弥梯罗MithilaCounty) BMithila弥梯罗() feud.Barony {
	return c.弥梯罗Mithila
}
    
func (c *弥梯罗MithilaCounty) BSaligrama舍利伽罗摩() feud.Barony {
	return c.舍利伽罗摩Saligrama
}
    
func (c *弥梯罗MithilaCounty) BSiwan锡万() feud.Barony {
	return c.锡万Siwan
}
    
func (c *弥梯罗MithilaCounty) BSugauna须乔那() feud.Barony {
	return c.须乔那Sugauna
}
    
func (c *弥梯罗MithilaCounty) BYavamajjhaka耶婆摩阇迦() feud.Barony {
	return c.耶婆摩阇迦Yavamajjhaka
}
    
var CMithila弥梯罗 MithilaCounty = &弥梯罗MithilaCounty{}

func init() {
	f := CMithila弥梯罗.(*弥梯罗MithilaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1419",
		Title:     "mithila",
		TitleName: "弥梯罗",
		TitleCode: "c_mithila",
		Baronies:  map[string]feud.Barony{},
	}

	f.达罗傍伽Darbhanga = BDarbhanga达罗傍伽
	f.达罗傍伽Darbhanga.SetParent(f)
	
	f.诃耆城Hajipur = BHajipur诃耆城
	f.诃耆城Hajipur.SetParent(f)
	
	f.弥梯罗Mithila = BMithila弥梯罗
	f.弥梯罗Mithila.SetParent(f)
	
	f.舍利伽罗摩Saligrama = BSaligrama舍利伽罗摩
	f.舍利伽罗摩Saligrama.SetParent(f)
	
	f.锡万Siwan = BSiwan锡万
	f.锡万Siwan.SetParent(f)
	
	f.须乔那Sugauna = BSugauna须乔那
	f.须乔那Sugauna.SetParent(f)
	
	f.耶婆摩阇迦Yavamajjhaka = BYavamajjhaka耶婆摩阇迦
	f.耶婆摩阇迦Yavamajjhaka.SetParent(f)
	
}
