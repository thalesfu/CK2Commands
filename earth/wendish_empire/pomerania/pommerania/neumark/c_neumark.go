package neumark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NeumarkCounty interface {
    feud.County
    BBernstein伯恩斯坦() 	feud.Barony
    BDramburg德兰堡() 	feud.Barony
    BDrawa德拉瓦() 	feud.Barony
    BDriesen德里森() 	feud.Barony
    BSantok桑托克() 	feud.Barony
    BSchivelbein席韦尔拜因() 	feud.Barony
    BTempelburg腾佩尔堡() 	feud.Barony
}

type 诺伊马尔克NeumarkCounty struct {
	feud.BaseCounty
	伯恩斯坦Bernstein 	feud.Barony
	德兰堡Dramburg 	feud.Barony
	德拉瓦Drawa 	feud.Barony
	德里森Driesen 	feud.Barony
	桑托克Santok 	feud.Barony
	席韦尔拜因Schivelbein 	feud.Barony
	腾佩尔堡Tempelburg 	feud.Barony
}

func (c *诺伊马尔克NeumarkCounty) BBernstein伯恩斯坦() feud.Barony {
	return c.伯恩斯坦Bernstein
}
    
func (c *诺伊马尔克NeumarkCounty) BDramburg德兰堡() feud.Barony {
	return c.德兰堡Dramburg
}
    
func (c *诺伊马尔克NeumarkCounty) BDrawa德拉瓦() feud.Barony {
	return c.德拉瓦Drawa
}
    
func (c *诺伊马尔克NeumarkCounty) BDriesen德里森() feud.Barony {
	return c.德里森Driesen
}
    
func (c *诺伊马尔克NeumarkCounty) BSantok桑托克() feud.Barony {
	return c.桑托克Santok
}
    
func (c *诺伊马尔克NeumarkCounty) BSchivelbein席韦尔拜因() feud.Barony {
	return c.席韦尔拜因Schivelbein
}
    
func (c *诺伊马尔克NeumarkCounty) BTempelburg腾佩尔堡() feud.Barony {
	return c.腾佩尔堡Tempelburg
}
    
var CNeumark诺伊马尔克 NeumarkCounty = &诺伊马尔克NeumarkCounty{}

func init() {
	f := CNeumark诺伊马尔克.(*诺伊马尔克NeumarkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1588",
		Title:     "neumark",
		TitleName: "诺伊马尔克",
		TitleCode: "c_neumark",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯恩斯坦Bernstein = BBernstein伯恩斯坦
	f.伯恩斯坦Bernstein.SetParent(f)
	
	f.德兰堡Dramburg = BDramburg德兰堡
	f.德兰堡Dramburg.SetParent(f)
	
	f.德拉瓦Drawa = BDrawa德拉瓦
	f.德拉瓦Drawa.SetParent(f)
	
	f.德里森Driesen = BDriesen德里森
	f.德里森Driesen.SetParent(f)
	
	f.桑托克Santok = BSantok桑托克
	f.桑托克Santok.SetParent(f)
	
	f.席韦尔拜因Schivelbein = BSchivelbein席韦尔拜因
	f.席韦尔拜因Schivelbein.SetParent(f)
	
	f.腾佩尔堡Tempelburg = BTempelburg腾佩尔堡
	f.腾佩尔堡Tempelburg.SetParent(f)
	
}
