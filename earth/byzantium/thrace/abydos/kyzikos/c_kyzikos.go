package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KyzikosCounty interface {
    feud.County
    BAdrianutherai阿德亚诺特莱() 	feud.Barony
    BArisbe阿里斯贝() 	feud.Barony
    BArtake阿尔代基() 	feud.Barony
    BKremasti克雷马斯蒂() 	feud.Barony
    BKyzikos基齐库斯() 	feud.Barony
    BMilitopolis米利波波利斯() 	feud.Barony
    BMyrina密里那() 	feud.Barony
    BPercote皮尔科特() 	feud.Barony
}

type 基齐库斯KyzikosCounty struct {
	feud.BaseCounty
	阿德亚诺特莱Adrianutherai 	feud.Barony
	阿里斯贝Arisbe 	feud.Barony
	阿尔代基Artake 	feud.Barony
	克雷马斯蒂Kremasti 	feud.Barony
	基齐库斯Kyzikos 	feud.Barony
	米利波波利斯Militopolis 	feud.Barony
	密里那Myrina 	feud.Barony
	皮尔科特Percote 	feud.Barony
}

func (c *基齐库斯KyzikosCounty) BAdrianutherai阿德亚诺特莱() feud.Barony {
	return c.阿德亚诺特莱Adrianutherai
}
    
func (c *基齐库斯KyzikosCounty) BArisbe阿里斯贝() feud.Barony {
	return c.阿里斯贝Arisbe
}
    
func (c *基齐库斯KyzikosCounty) BArtake阿尔代基() feud.Barony {
	return c.阿尔代基Artake
}
    
func (c *基齐库斯KyzikosCounty) BKremasti克雷马斯蒂() feud.Barony {
	return c.克雷马斯蒂Kremasti
}
    
func (c *基齐库斯KyzikosCounty) BKyzikos基齐库斯() feud.Barony {
	return c.基齐库斯Kyzikos
}
    
func (c *基齐库斯KyzikosCounty) BMilitopolis米利波波利斯() feud.Barony {
	return c.米利波波利斯Militopolis
}
    
func (c *基齐库斯KyzikosCounty) BMyrina密里那() feud.Barony {
	return c.密里那Myrina
}
    
func (c *基齐库斯KyzikosCounty) BPercote皮尔科特() feud.Barony {
	return c.皮尔科特Percote
}
    
var CKyzikos基齐库斯 KyzikosCounty = &基齐库斯KyzikosCounty{}

func init() {
	f := CKyzikos基齐库斯.(*基齐库斯KyzikosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "743",
		Title:     "kyzikos",
		TitleName: "基齐库斯",
		TitleCode: "c_kyzikos",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德亚诺特莱Adrianutherai = BAdrianutherai阿德亚诺特莱
	f.阿德亚诺特莱Adrianutherai.SetParent(f)
	
	f.阿里斯贝Arisbe = BArisbe阿里斯贝
	f.阿里斯贝Arisbe.SetParent(f)
	
	f.阿尔代基Artake = BArtake阿尔代基
	f.阿尔代基Artake.SetParent(f)
	
	f.克雷马斯蒂Kremasti = BKremasti克雷马斯蒂
	f.克雷马斯蒂Kremasti.SetParent(f)
	
	f.基齐库斯Kyzikos = BKyzikos基齐库斯
	f.基齐库斯Kyzikos.SetParent(f)
	
	f.米利波波利斯Militopolis = BMilitopolis米利波波利斯
	f.米利波波利斯Militopolis.SetParent(f)
	
	f.密里那Myrina = BMyrina密里那
	f.密里那Myrina.SetParent(f)
	
	f.皮尔科特Percote = BPercote皮尔科特
	f.皮尔科特Percote.SetParent(f)
	
}
