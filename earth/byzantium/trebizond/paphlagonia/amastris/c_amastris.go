package amastris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmastrisCounty interface {
    feud.County
    BAmasra阿马斯拉() 	feud.Barony
    BAmastris阿马斯特里斯() 	feud.Barony
    BCromna克罗姆纳() 	feud.Barony
    BCyttoros库托洛斯() 	feud.Barony
    BSesamus塞萨摩斯() 	feud.Barony
    BTieion提翁() 	feud.Barony
    BTium提乌姆() 	feud.Barony
}

type 阿马斯特里斯AmastrisCounty struct {
	feud.BaseCounty
	阿马斯拉Amasra 	feud.Barony
	阿马斯特里斯Amastris 	feud.Barony
	克罗姆纳Cromna 	feud.Barony
	库托洛斯Cyttoros 	feud.Barony
	塞萨摩斯Sesamus 	feud.Barony
	提翁Tieion 	feud.Barony
	提乌姆Tium 	feud.Barony
}

func (c *阿马斯特里斯AmastrisCounty) BAmasra阿马斯拉() feud.Barony {
	return c.阿马斯拉Amasra
}
    
func (c *阿马斯特里斯AmastrisCounty) BAmastris阿马斯特里斯() feud.Barony {
	return c.阿马斯特里斯Amastris
}
    
func (c *阿马斯特里斯AmastrisCounty) BCromna克罗姆纳() feud.Barony {
	return c.克罗姆纳Cromna
}
    
func (c *阿马斯特里斯AmastrisCounty) BCyttoros库托洛斯() feud.Barony {
	return c.库托洛斯Cyttoros
}
    
func (c *阿马斯特里斯AmastrisCounty) BSesamus塞萨摩斯() feud.Barony {
	return c.塞萨摩斯Sesamus
}
    
func (c *阿马斯特里斯AmastrisCounty) BTieion提翁() feud.Barony {
	return c.提翁Tieion
}
    
func (c *阿马斯特里斯AmastrisCounty) BTium提乌姆() feud.Barony {
	return c.提乌姆Tium
}
    
var CAmastris阿马斯特里斯 AmastrisCounty = &阿马斯特里斯AmastrisCounty{}

func init() {
	f := CAmastris阿马斯特里斯.(*阿马斯特里斯AmastrisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1928",
		Title:     "amastris",
		TitleName: "阿马斯特里斯",
		TitleCode: "c_amastris",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿马斯拉Amasra = BAmasra阿马斯拉
	f.阿马斯拉Amasra.SetParent(f)
	
	f.阿马斯特里斯Amastris = BAmastris阿马斯特里斯
	f.阿马斯特里斯Amastris.SetParent(f)
	
	f.克罗姆纳Cromna = BCromna克罗姆纳
	f.克罗姆纳Cromna.SetParent(f)
	
	f.库托洛斯Cyttoros = BCyttoros库托洛斯
	f.库托洛斯Cyttoros.SetParent(f)
	
	f.塞萨摩斯Sesamus = BSesamus塞萨摩斯
	f.塞萨摩斯Sesamus.SetParent(f)
	
	f.提翁Tieion = BTieion提翁
	f.提翁Tieion.SetParent(f)
	
	f.提乌姆Tium = BTium提乌姆
	f.提乌姆Tium.SetParent(f)
	
}
