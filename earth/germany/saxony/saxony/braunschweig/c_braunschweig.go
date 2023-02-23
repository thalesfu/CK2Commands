package braunschweig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BraunschweigCounty interface {
	feud.County
	BBielefeld比勒费尔德() feud.Barony
	BBraunschweig布伦瑞克() feud.Barony
	BGandersheim甘德斯海姆() feud.Barony
	BHelmstedt黑尔姆施泰特() feud.Barony
	BHildesheim希尔德斯海姆() feud.Barony
	BWaldeck瓦尔德克() feud.Barony
	BWolfenbuttel沃尔芬比特尔() feud.Barony
}

type 布伦瑞克BraunschweigCounty struct {
	feud.BaseCounty
	比勒费尔德Bielefeld     feud.Barony
	布伦瑞克Braunschweig   feud.Barony
	甘德斯海姆Gandersheim   feud.Barony
	黑尔姆施泰特Helmstedt    feud.Barony
	希尔德斯海姆Hildesheim   feud.Barony
	瓦尔德克Waldeck        feud.Barony
	沃尔芬比特尔Wolfenbuttel feud.Barony
}

func (c *布伦瑞克BraunschweigCounty) BBielefeld比勒费尔德() feud.Barony {
	return c.比勒费尔德Bielefeld
}

func (c *布伦瑞克BraunschweigCounty) BBraunschweig布伦瑞克() feud.Barony {
	return c.布伦瑞克Braunschweig
}

func (c *布伦瑞克BraunschweigCounty) BGandersheim甘德斯海姆() feud.Barony {
	return c.甘德斯海姆Gandersheim
}

func (c *布伦瑞克BraunschweigCounty) BHelmstedt黑尔姆施泰特() feud.Barony {
	return c.黑尔姆施泰特Helmstedt
}

func (c *布伦瑞克BraunschweigCounty) BHildesheim希尔德斯海姆() feud.Barony {
	return c.希尔德斯海姆Hildesheim
}

func (c *布伦瑞克BraunschweigCounty) BWaldeck瓦尔德克() feud.Barony {
	return c.瓦尔德克Waldeck
}

func (c *布伦瑞克BraunschweigCounty) BWolfenbuttel沃尔芬比特尔() feud.Barony {
	return c.沃尔芬比特尔Wolfenbuttel
}

var CBraunschweig布伦瑞克 BraunschweigCounty = &布伦瑞克BraunschweigCounty{}

func init() {
	f := CBraunschweig布伦瑞克.(*布伦瑞克BraunschweigCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "257",
		Title:     "braunschweig",
		TitleName: "布伦瑞克",
		TitleCode: "c_braunschweig",
		Baronies:  map[string]feud.Barony{},
	}

	f.比勒费尔德Bielefeld = BBielefeld比勒费尔德
	f.比勒费尔德Bielefeld.SetParent(f)

	f.布伦瑞克Braunschweig = BBraunschweig布伦瑞克
	f.布伦瑞克Braunschweig.SetParent(f)

	f.甘德斯海姆Gandersheim = BGandersheim甘德斯海姆
	f.甘德斯海姆Gandersheim.SetParent(f)

	f.黑尔姆施泰特Helmstedt = BHelmstedt黑尔姆施泰特
	f.黑尔姆施泰特Helmstedt.SetParent(f)

	f.希尔德斯海姆Hildesheim = BHildesheim希尔德斯海姆
	f.希尔德斯海姆Hildesheim.SetParent(f)

	f.瓦尔德克Waldeck = BWaldeck瓦尔德克
	f.瓦尔德克Waldeck.SetParent(f)

	f.沃尔芬比特尔Wolfenbuttel = BWolfenbuttel沃尔芬比特尔
	f.沃尔芬比特尔Wolfenbuttel.SetParent(f)

}
