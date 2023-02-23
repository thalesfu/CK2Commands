package rzheva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RzhevaCounty interface {
	feud.County
	BDeshevki捷舍夫基() feud.Barony
	BGorodok戈罗多克() feud.Barony
	BMednoye梅德诺耶() feud.Barony
	BPleshki普列什基() feud.Barony
	BRzheva勒热瓦() feud.Barony
	BVysokoye维索科耶() feud.Barony
	BZelenyy泽廖内() feud.Barony
}

type 勒热瓦RzhevaCounty struct {
	feud.BaseCounty
	捷舍夫基Deshevki feud.Barony
	戈罗多克Gorodok  feud.Barony
	梅德诺耶Mednoye  feud.Barony
	普列什基Pleshki  feud.Barony
	勒热瓦Rzheva    feud.Barony
	维索科耶Vysokoye feud.Barony
	泽廖内Zelenyy   feud.Barony
}

func (c *勒热瓦RzhevaCounty) BDeshevki捷舍夫基() feud.Barony {
	return c.捷舍夫基Deshevki
}

func (c *勒热瓦RzhevaCounty) BGorodok戈罗多克() feud.Barony {
	return c.戈罗多克Gorodok
}

func (c *勒热瓦RzhevaCounty) BMednoye梅德诺耶() feud.Barony {
	return c.梅德诺耶Mednoye
}

func (c *勒热瓦RzhevaCounty) BPleshki普列什基() feud.Barony {
	return c.普列什基Pleshki
}

func (c *勒热瓦RzhevaCounty) BRzheva勒热瓦() feud.Barony {
	return c.勒热瓦Rzheva
}

func (c *勒热瓦RzhevaCounty) BVysokoye维索科耶() feud.Barony {
	return c.维索科耶Vysokoye
}

func (c *勒热瓦RzhevaCounty) BZelenyy泽廖内() feud.Barony {
	return c.泽廖内Zelenyy
}

var CRzheva勒热瓦 RzhevaCounty = &勒热瓦RzhevaCounty{}

func init() {
	f := CRzheva勒热瓦.(*勒热瓦RzhevaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1668",
		Title:     "rzheva",
		TitleName: "勒热瓦",
		TitleCode: "c_rzheva",
		Baronies:  map[string]feud.Barony{},
	}

	f.捷舍夫基Deshevki = BDeshevki捷舍夫基
	f.捷舍夫基Deshevki.SetParent(f)

	f.戈罗多克Gorodok = BGorodok戈罗多克
	f.戈罗多克Gorodok.SetParent(f)

	f.梅德诺耶Mednoye = BMednoye梅德诺耶
	f.梅德诺耶Mednoye.SetParent(f)

	f.普列什基Pleshki = BPleshki普列什基
	f.普列什基Pleshki.SetParent(f)

	f.勒热瓦Rzheva = BRzheva勒热瓦
	f.勒热瓦Rzheva.SetParent(f)

	f.维索科耶Vysokoye = BVysokoye维索科耶
	f.维索科耶Vysokoye.SetParent(f)

	f.泽廖内Zelenyy = BZelenyy泽廖内
	f.泽廖内Zelenyy.SetParent(f)

}
