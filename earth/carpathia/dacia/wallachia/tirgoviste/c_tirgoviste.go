package tirgoviste

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TirgovisteCounty interface {
    feud.County
    BBucov布科夫() 	feud.Barony
    BBuzau布泽乌() 	feud.Barony
    BRamnicu_sarat勒姆尼库瑟拉特() 	feud.Barony
    BSnagov斯纳戈夫() 	feud.Barony
    BTabla_butii普楚尔布蒂() 	feud.Barony
    BTargoviste特尔戈维什泰() 	feud.Barony
    BValeni沃莱尼() 	feud.Barony
}

type 特尔戈维什泰TirgovisteCounty struct {
	feud.BaseCounty
	布科夫Bucov 	feud.Barony
	布泽乌Buzau 	feud.Barony
	勒姆尼库瑟拉特Ramnicu_sarat 	feud.Barony
	斯纳戈夫Snagov 	feud.Barony
	普楚尔布蒂Tabla_butii 	feud.Barony
	特尔戈维什泰Targoviste 	feud.Barony
	沃莱尼Valeni 	feud.Barony
}

func (c *特尔戈维什泰TirgovisteCounty) BBucov布科夫() feud.Barony {
	return c.布科夫Bucov
}
    
func (c *特尔戈维什泰TirgovisteCounty) BBuzau布泽乌() feud.Barony {
	return c.布泽乌Buzau
}
    
func (c *特尔戈维什泰TirgovisteCounty) BRamnicu_sarat勒姆尼库瑟拉特() feud.Barony {
	return c.勒姆尼库瑟拉特Ramnicu_sarat
}
    
func (c *特尔戈维什泰TirgovisteCounty) BSnagov斯纳戈夫() feud.Barony {
	return c.斯纳戈夫Snagov
}
    
func (c *特尔戈维什泰TirgovisteCounty) BTabla_butii普楚尔布蒂() feud.Barony {
	return c.普楚尔布蒂Tabla_butii
}
    
func (c *特尔戈维什泰TirgovisteCounty) BTargoviste特尔戈维什泰() feud.Barony {
	return c.特尔戈维什泰Targoviste
}
    
func (c *特尔戈维什泰TirgovisteCounty) BValeni沃莱尼() feud.Barony {
	return c.沃莱尼Valeni
}
    
var CTirgoviste特尔戈维什泰 TirgovisteCounty = &特尔戈维什泰TirgovisteCounty{}

func init() {
	f := CTirgoviste特尔戈维什泰.(*特尔戈维什泰TirgovisteCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "515",
		Title:     "tirgoviste",
		TitleName: "特尔戈维什泰",
		TitleCode: "c_tirgoviste",
		Baronies:  map[string]feud.Barony{},
	}

	f.布科夫Bucov = BBucov布科夫
	f.布科夫Bucov.SetParent(f)
	
	f.布泽乌Buzau = BBuzau布泽乌
	f.布泽乌Buzau.SetParent(f)
	
	f.勒姆尼库瑟拉特Ramnicu_sarat = BRamnicu_sarat勒姆尼库瑟拉特
	f.勒姆尼库瑟拉特Ramnicu_sarat.SetParent(f)
	
	f.斯纳戈夫Snagov = BSnagov斯纳戈夫
	f.斯纳戈夫Snagov.SetParent(f)
	
	f.普楚尔布蒂Tabla_butii = BTabla_butii普楚尔布蒂
	f.普楚尔布蒂Tabla_butii.SetParent(f)
	
	f.特尔戈维什泰Targoviste = BTargoviste特尔戈维什泰
	f.特尔戈维什泰Targoviste.SetParent(f)
	
	f.沃莱尼Valeni = BValeni沃莱尼
	f.沃莱尼Valeni.SetParent(f)
	
}
