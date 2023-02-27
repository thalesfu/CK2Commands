package uvek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UvekCounty interface {
    feud.County
    BBurnyy布尔内() 	feud.Barony
    BLipovka利波夫卡() 	feud.Barony
    BMarks马克思() 	feud.Barony
    BShumeyka舒梅卡() 	feud.Barony
    BSmelovka斯梅洛夫卡() 	feud.Barony
    BTitorenko季托连科() 	feud.Barony
    BUvek乌韦克() 	feud.Barony
}

type 乌韦克UvekCounty struct {
	feud.BaseCounty
	布尔内Burnyy 	feud.Barony
	利波夫卡Lipovka 	feud.Barony
	马克思Marks 	feud.Barony
	舒梅卡Shumeyka 	feud.Barony
	斯梅洛夫卡Smelovka 	feud.Barony
	季托连科Titorenko 	feud.Barony
	乌韦克Uvek 	feud.Barony
}

func (c *乌韦克UvekCounty) BBurnyy布尔内() feud.Barony {
	return c.布尔内Burnyy
}
    
func (c *乌韦克UvekCounty) BLipovka利波夫卡() feud.Barony {
	return c.利波夫卡Lipovka
}
    
func (c *乌韦克UvekCounty) BMarks马克思() feud.Barony {
	return c.马克思Marks
}
    
func (c *乌韦克UvekCounty) BShumeyka舒梅卡() feud.Barony {
	return c.舒梅卡Shumeyka
}
    
func (c *乌韦克UvekCounty) BSmelovka斯梅洛夫卡() feud.Barony {
	return c.斯梅洛夫卡Smelovka
}
    
func (c *乌韦克UvekCounty) BTitorenko季托连科() feud.Barony {
	return c.季托连科Titorenko
}
    
func (c *乌韦克UvekCounty) BUvek乌韦克() feud.Barony {
	return c.乌韦克Uvek
}
    
var CUvek乌韦克 UvekCounty = &乌韦克UvekCounty{}

func init() {
	f := CUvek乌韦克.(*乌韦克UvekCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1814",
		Title:     "uvek",
		TitleName: "乌韦克",
		TitleCode: "c_uvek",
		Baronies:  map[string]feud.Barony{},
	}

	f.布尔内Burnyy = BBurnyy布尔内
	f.布尔内Burnyy.SetParent(f)
	
	f.利波夫卡Lipovka = BLipovka利波夫卡
	f.利波夫卡Lipovka.SetParent(f)
	
	f.马克思Marks = BMarks马克思
	f.马克思Marks.SetParent(f)
	
	f.舒梅卡Shumeyka = BShumeyka舒梅卡
	f.舒梅卡Shumeyka.SetParent(f)
	
	f.斯梅洛夫卡Smelovka = BSmelovka斯梅洛夫卡
	f.斯梅洛夫卡Smelovka.SetParent(f)
	
	f.季托连科Titorenko = BTitorenko季托连科
	f.季托连科Titorenko.SetParent(f)
	
	f.乌韦克Uvek = BUvek乌韦克
	f.乌韦克Uvek.SetParent(f)
	
}
