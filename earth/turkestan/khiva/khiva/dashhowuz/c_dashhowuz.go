package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DashhowuzCounty interface {
    feud.County
    BAkdepe阿克杰佩() 	feud.Barony
    BBoldumsaz博尔杜姆萨兹() 	feud.Barony
    BDashhowuz达什霍武兹() 	feud.Barony
    BGorogly吉奥罗格雷() 	feud.Barony
    BGubadag库巴达格() 	feud.Barony
    BKunyaurgench古玉龙杰赤() 	feud.Barony
    BTagta塔赫塔() 	feud.Barony
    BYylanly伊厄兰雷() 	feud.Barony
}

type 达什霍武兹DashhowuzCounty struct {
	feud.BaseCounty
	阿克杰佩Akdepe 	feud.Barony
	博尔杜姆萨兹Boldumsaz 	feud.Barony
	达什霍武兹Dashhowuz 	feud.Barony
	吉奥罗格雷Gorogly 	feud.Barony
	库巴达格Gubadag 	feud.Barony
	古玉龙杰赤Kunyaurgench 	feud.Barony
	塔赫塔Tagta 	feud.Barony
	伊厄兰雷Yylanly 	feud.Barony
}

func (c *达什霍武兹DashhowuzCounty) BAkdepe阿克杰佩() feud.Barony {
	return c.阿克杰佩Akdepe
}
    
func (c *达什霍武兹DashhowuzCounty) BBoldumsaz博尔杜姆萨兹() feud.Barony {
	return c.博尔杜姆萨兹Boldumsaz
}
    
func (c *达什霍武兹DashhowuzCounty) BDashhowuz达什霍武兹() feud.Barony {
	return c.达什霍武兹Dashhowuz
}
    
func (c *达什霍武兹DashhowuzCounty) BGorogly吉奥罗格雷() feud.Barony {
	return c.吉奥罗格雷Gorogly
}
    
func (c *达什霍武兹DashhowuzCounty) BGubadag库巴达格() feud.Barony {
	return c.库巴达格Gubadag
}
    
func (c *达什霍武兹DashhowuzCounty) BKunyaurgench古玉龙杰赤() feud.Barony {
	return c.古玉龙杰赤Kunyaurgench
}
    
func (c *达什霍武兹DashhowuzCounty) BTagta塔赫塔() feud.Barony {
	return c.塔赫塔Tagta
}
    
func (c *达什霍武兹DashhowuzCounty) BYylanly伊厄兰雷() feud.Barony {
	return c.伊厄兰雷Yylanly
}
    
var CDashhowuz达什霍武兹 DashhowuzCounty = &达什霍武兹DashhowuzCounty{}

func init() {
	f := CDashhowuz达什霍武兹.(*达什霍武兹DashhowuzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "902",
		Title:     "dashhowuz",
		TitleName: "达什霍武兹",
		TitleCode: "c_dashhowuz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克杰佩Akdepe = BAkdepe阿克杰佩
	f.阿克杰佩Akdepe.SetParent(f)
	
	f.博尔杜姆萨兹Boldumsaz = BBoldumsaz博尔杜姆萨兹
	f.博尔杜姆萨兹Boldumsaz.SetParent(f)
	
	f.达什霍武兹Dashhowuz = BDashhowuz达什霍武兹
	f.达什霍武兹Dashhowuz.SetParent(f)
	
	f.吉奥罗格雷Gorogly = BGorogly吉奥罗格雷
	f.吉奥罗格雷Gorogly.SetParent(f)
	
	f.库巴达格Gubadag = BGubadag库巴达格
	f.库巴达格Gubadag.SetParent(f)
	
	f.古玉龙杰赤Kunyaurgench = BKunyaurgench古玉龙杰赤
	f.古玉龙杰赤Kunyaurgench.SetParent(f)
	
	f.塔赫塔Tagta = BTagta塔赫塔
	f.塔赫塔Tagta.SetParent(f)
	
	f.伊厄兰雷Yylanly = BYylanly伊厄兰雷
	f.伊厄兰雷Yylanly.SetParent(f)
	
}
