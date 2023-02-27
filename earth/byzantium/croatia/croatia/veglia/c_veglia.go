package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VegliaCounty interface {
    feud.County
    BBakar巴卡尔() 	feud.Barony
    BCres茨雷斯() 	feud.Barony
    BCrikvenica茨里克韦尼察() 	feud.Barony
    BFrankopan弗兰科潘() 	feud.Barony
    BKraljevica克拉列维察() 	feud.Barony
    BKrk克尔克() 	feud.Barony
    BVeglia韦利亚() 	feud.Barony
    BVrbovsko弗尔博夫斯科() 	feud.Barony
}

type 韦利亚VegliaCounty struct {
	feud.BaseCounty
	巴卡尔Bakar 	feud.Barony
	茨雷斯Cres 	feud.Barony
	茨里克韦尼察Crikvenica 	feud.Barony
	弗兰科潘Frankopan 	feud.Barony
	克拉列维察Kraljevica 	feud.Barony
	克尔克Krk 	feud.Barony
	韦利亚Veglia 	feud.Barony
	弗尔博夫斯科Vrbovsko 	feud.Barony
}

func (c *韦利亚VegliaCounty) BBakar巴卡尔() feud.Barony {
	return c.巴卡尔Bakar
}
    
func (c *韦利亚VegliaCounty) BCres茨雷斯() feud.Barony {
	return c.茨雷斯Cres
}
    
func (c *韦利亚VegliaCounty) BCrikvenica茨里克韦尼察() feud.Barony {
	return c.茨里克韦尼察Crikvenica
}
    
func (c *韦利亚VegliaCounty) BFrankopan弗兰科潘() feud.Barony {
	return c.弗兰科潘Frankopan
}
    
func (c *韦利亚VegliaCounty) BKraljevica克拉列维察() feud.Barony {
	return c.克拉列维察Kraljevica
}
    
func (c *韦利亚VegliaCounty) BKrk克尔克() feud.Barony {
	return c.克尔克Krk
}
    
func (c *韦利亚VegliaCounty) BVeglia韦利亚() feud.Barony {
	return c.韦利亚Veglia
}
    
func (c *韦利亚VegliaCounty) BVrbovsko弗尔博夫斯科() feud.Barony {
	return c.弗尔博夫斯科Vrbovsko
}
    
var CVeglia韦利亚 VegliaCounty = &韦利亚VegliaCounty{}

func init() {
	f := CVeglia韦利亚.(*韦利亚VegliaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "459",
		Title:     "veglia",
		TitleName: "韦利亚",
		TitleCode: "c_veglia",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴卡尔Bakar = BBakar巴卡尔
	f.巴卡尔Bakar.SetParent(f)
	
	f.茨雷斯Cres = BCres茨雷斯
	f.茨雷斯Cres.SetParent(f)
	
	f.茨里克韦尼察Crikvenica = BCrikvenica茨里克韦尼察
	f.茨里克韦尼察Crikvenica.SetParent(f)
	
	f.弗兰科潘Frankopan = BFrankopan弗兰科潘
	f.弗兰科潘Frankopan.SetParent(f)
	
	f.克拉列维察Kraljevica = BKraljevica克拉列维察
	f.克拉列维察Kraljevica.SetParent(f)
	
	f.克尔克Krk = BKrk克尔克
	f.克尔克Krk.SetParent(f)
	
	f.韦利亚Veglia = BVeglia韦利亚
	f.韦利亚Veglia.SetParent(f)
	
	f.弗尔博夫斯科Vrbovsko = BVrbovsko弗尔博夫斯科
	f.弗尔博夫斯科Vrbovsko.SetParent(f)
	
}
