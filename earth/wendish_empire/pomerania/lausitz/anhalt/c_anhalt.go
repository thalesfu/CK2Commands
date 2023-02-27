package anhalt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnhaltCounty interface {
    feud.County
    BCoswig科斯维希() 	feud.Barony
    BDessau德绍() 	feud.Barony
    BDornburg_zeist多恩堡() 	feud.Barony
    BLindau_zeist林道() 	feud.Barony
    BWittenberg_zeist维滕贝格() 	feud.Barony
    BZerbst采尔布斯特() 	feud.Barony
}

type 采尔布斯特AnhaltCounty struct {
	feud.BaseCounty
	科斯维希Coswig 	feud.Barony
	德绍Dessau 	feud.Barony
	多恩堡Dornburg_zeist 	feud.Barony
	林道Lindau_zeist 	feud.Barony
	维滕贝格Wittenberg_zeist 	feud.Barony
	采尔布斯特Zerbst 	feud.Barony
}

func (c *采尔布斯特AnhaltCounty) BCoswig科斯维希() feud.Barony {
	return c.科斯维希Coswig
}
    
func (c *采尔布斯特AnhaltCounty) BDessau德绍() feud.Barony {
	return c.德绍Dessau
}
    
func (c *采尔布斯特AnhaltCounty) BDornburg_zeist多恩堡() feud.Barony {
	return c.多恩堡Dornburg_zeist
}
    
func (c *采尔布斯特AnhaltCounty) BLindau_zeist林道() feud.Barony {
	return c.林道Lindau_zeist
}
    
func (c *采尔布斯特AnhaltCounty) BWittenberg_zeist维滕贝格() feud.Barony {
	return c.维滕贝格Wittenberg_zeist
}
    
func (c *采尔布斯特AnhaltCounty) BZerbst采尔布斯特() feud.Barony {
	return c.采尔布斯特Zerbst
}
    
var CAnhalt采尔布斯特 AnhaltCounty = &采尔布斯特AnhaltCounty{}

func init() {
	f := CAnhalt采尔布斯特.(*采尔布斯特AnhaltCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "310",
		Title:     "anhalt",
		TitleName: "采尔布斯特",
		TitleCode: "c_anhalt",
		Baronies:  map[string]feud.Barony{},
	}

	f.科斯维希Coswig = BCoswig科斯维希
	f.科斯维希Coswig.SetParent(f)
	
	f.德绍Dessau = BDessau德绍
	f.德绍Dessau.SetParent(f)
	
	f.多恩堡Dornburg_zeist = BDornburg_zeist多恩堡
	f.多恩堡Dornburg_zeist.SetParent(f)
	
	f.林道Lindau_zeist = BLindau_zeist林道
	f.林道Lindau_zeist.SetParent(f)
	
	f.维滕贝格Wittenberg_zeist = BWittenberg_zeist维滕贝格
	f.维滕贝格Wittenberg_zeist.SetParent(f)
	
	f.采尔布斯特Zerbst = BZerbst采尔布斯特
	f.采尔布斯特Zerbst.SetParent(f)
	
}
