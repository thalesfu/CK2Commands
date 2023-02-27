package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KumaCounty interface {
    feud.County
    BArslanbek阿尔斯兰别克() 	feud.Barony
    BBabayurt巴巴尤尔特() 	feud.Barony
    BDylym德雷姆() 	feud.Barony
    BGums居姆谢() 	feud.Barony
    BKizilyurt克孜勒尤尔特() 	feud.Barony
    BKizlyar基兹利亚尔() 	feud.Barony
    BKuliyurt库里_尤尔特() 	feud.Barony
    BTereklimekteb捷列克利_梅克捷布() 	feud.Barony
}

type 库马KumaCounty struct {
	feud.BaseCounty
	阿尔斯兰别克Arslanbek 	feud.Barony
	巴巴尤尔特Babayurt 	feud.Barony
	德雷姆Dylym 	feud.Barony
	居姆谢Gums 	feud.Barony
	克孜勒尤尔特Kizilyurt 	feud.Barony
	基兹利亚尔Kizlyar 	feud.Barony
	库里_尤尔特Kuliyurt 	feud.Barony
	捷列克利_梅克捷布Tereklimekteb 	feud.Barony
}

func (c *库马KumaCounty) BArslanbek阿尔斯兰别克() feud.Barony {
	return c.阿尔斯兰别克Arslanbek
}
    
func (c *库马KumaCounty) BBabayurt巴巴尤尔特() feud.Barony {
	return c.巴巴尤尔特Babayurt
}
    
func (c *库马KumaCounty) BDylym德雷姆() feud.Barony {
	return c.德雷姆Dylym
}
    
func (c *库马KumaCounty) BGums居姆谢() feud.Barony {
	return c.居姆谢Gums
}
    
func (c *库马KumaCounty) BKizilyurt克孜勒尤尔特() feud.Barony {
	return c.克孜勒尤尔特Kizilyurt
}
    
func (c *库马KumaCounty) BKizlyar基兹利亚尔() feud.Barony {
	return c.基兹利亚尔Kizlyar
}
    
func (c *库马KumaCounty) BKuliyurt库里_尤尔特() feud.Barony {
	return c.库里_尤尔特Kuliyurt
}
    
func (c *库马KumaCounty) BTereklimekteb捷列克利_梅克捷布() feud.Barony {
	return c.捷列克利_梅克捷布Tereklimekteb
}
    
var CKuma库马 KumaCounty = &库马KumaCounty{}

func init() {
	f := CKuma库马.(*库马KumaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "604",
		Title:     "kuma",
		TitleName: "库马",
		TitleCode: "c_kuma",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔斯兰别克Arslanbek = BArslanbek阿尔斯兰别克
	f.阿尔斯兰别克Arslanbek.SetParent(f)
	
	f.巴巴尤尔特Babayurt = BBabayurt巴巴尤尔特
	f.巴巴尤尔特Babayurt.SetParent(f)
	
	f.德雷姆Dylym = BDylym德雷姆
	f.德雷姆Dylym.SetParent(f)
	
	f.居姆谢Gums = BGums居姆谢
	f.居姆谢Gums.SetParent(f)
	
	f.克孜勒尤尔特Kizilyurt = BKizilyurt克孜勒尤尔特
	f.克孜勒尤尔特Kizilyurt.SetParent(f)
	
	f.基兹利亚尔Kizlyar = BKizlyar基兹利亚尔
	f.基兹利亚尔Kizlyar.SetParent(f)
	
	f.库里_尤尔特Kuliyurt = BKuliyurt库里_尤尔特
	f.库里_尤尔特Kuliyurt.SetParent(f)
	
	f.捷列克利_梅克捷布Tereklimekteb = BTereklimekteb捷列克利_梅克捷布
	f.捷列克利_梅克捷布Tereklimekteb.SetParent(f)
	
}
