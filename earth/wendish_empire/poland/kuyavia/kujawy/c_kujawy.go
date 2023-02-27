package kujawy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KujawyCounty interface {
    feud.County
    BBrzesckujawskie库亚维地区布热希奇() 	feud.Barony
    BBygdoszcz比得哥什() 	feud.Barony
    BInowroclaw伊诺弗罗茨瓦夫() 	feud.Barony
    BKlodawa克沃达瓦() 	feud.Barony
    BKruszwica克鲁什维察() 	feud.Barony
    BLeczyca文奇察() 	feud.Barony
    BNieszawa涅沙瓦() 	feud.Barony
}

type 库亚维KujawyCounty struct {
	feud.BaseCounty
	库亚维地区布热希奇Brzesckujawskie 	feud.Barony
	比得哥什Bygdoszcz 	feud.Barony
	伊诺弗罗茨瓦夫Inowroclaw 	feud.Barony
	克沃达瓦Klodawa 	feud.Barony
	克鲁什维察Kruszwica 	feud.Barony
	文奇察Leczyca 	feud.Barony
	涅沙瓦Nieszawa 	feud.Barony
}

func (c *库亚维KujawyCounty) BBrzesckujawskie库亚维地区布热希奇() feud.Barony {
	return c.库亚维地区布热希奇Brzesckujawskie
}
    
func (c *库亚维KujawyCounty) BBygdoszcz比得哥什() feud.Barony {
	return c.比得哥什Bygdoszcz
}
    
func (c *库亚维KujawyCounty) BInowroclaw伊诺弗罗茨瓦夫() feud.Barony {
	return c.伊诺弗罗茨瓦夫Inowroclaw
}
    
func (c *库亚维KujawyCounty) BKlodawa克沃达瓦() feud.Barony {
	return c.克沃达瓦Klodawa
}
    
func (c *库亚维KujawyCounty) BKruszwica克鲁什维察() feud.Barony {
	return c.克鲁什维察Kruszwica
}
    
func (c *库亚维KujawyCounty) BLeczyca文奇察() feud.Barony {
	return c.文奇察Leczyca
}
    
func (c *库亚维KujawyCounty) BNieszawa涅沙瓦() feud.Barony {
	return c.涅沙瓦Nieszawa
}
    
var CKujawy库亚维 KujawyCounty = &库亚维KujawyCounty{}

func init() {
	f := CKujawy库亚维.(*库亚维KujawyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "428",
		Title:     "kujawy",
		TitleName: "库亚维",
		TitleCode: "c_kujawy",
		Baronies:  map[string]feud.Barony{},
	}

	f.库亚维地区布热希奇Brzesckujawskie = BBrzesckujawskie库亚维地区布热希奇
	f.库亚维地区布热希奇Brzesckujawskie.SetParent(f)
	
	f.比得哥什Bygdoszcz = BBygdoszcz比得哥什
	f.比得哥什Bygdoszcz.SetParent(f)
	
	f.伊诺弗罗茨瓦夫Inowroclaw = BInowroclaw伊诺弗罗茨瓦夫
	f.伊诺弗罗茨瓦夫Inowroclaw.SetParent(f)
	
	f.克沃达瓦Klodawa = BKlodawa克沃达瓦
	f.克沃达瓦Klodawa.SetParent(f)
	
	f.克鲁什维察Kruszwica = BKruszwica克鲁什维察
	f.克鲁什维察Kruszwica.SetParent(f)
	
	f.文奇察Leczyca = BLeczyca文奇察
	f.文奇察Leczyca.SetParent(f)
	
	f.涅沙瓦Nieszawa = BNieszawa涅沙瓦
	f.涅沙瓦Nieszawa.SetParent(f)
	
}
