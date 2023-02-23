package zamindawar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZamindawarCounty interface {
	feud.County
	BAfgilai阿夫吉莱() feud.Barony
	BAmandawar阿曼达瓦尔() feud.Barony
	BBishlang比什朗格() feud.Barony
	BKhalai哈莱() feud.Barony
	BLangi兰吉() feud.Barony
	BYmanda伊曼达() feud.Barony
	BZamindawar扎敏达瓦尔() feud.Barony
}

type 扎敏达瓦尔ZamindawarCounty struct {
	feud.BaseCounty
	阿夫吉莱Afgilai     feud.Barony
	阿曼达瓦尔Amandawar  feud.Barony
	比什朗格Bishlang    feud.Barony
	哈莱Khalai        feud.Barony
	兰吉Langi         feud.Barony
	伊曼达Ymanda       feud.Barony
	扎敏达瓦尔Zamindawar feud.Barony
}

func (c *扎敏达瓦尔ZamindawarCounty) BAfgilai阿夫吉莱() feud.Barony {
	return c.阿夫吉莱Afgilai
}

func (c *扎敏达瓦尔ZamindawarCounty) BAmandawar阿曼达瓦尔() feud.Barony {
	return c.阿曼达瓦尔Amandawar
}

func (c *扎敏达瓦尔ZamindawarCounty) BBishlang比什朗格() feud.Barony {
	return c.比什朗格Bishlang
}

func (c *扎敏达瓦尔ZamindawarCounty) BKhalai哈莱() feud.Barony {
	return c.哈莱Khalai
}

func (c *扎敏达瓦尔ZamindawarCounty) BLangi兰吉() feud.Barony {
	return c.兰吉Langi
}

func (c *扎敏达瓦尔ZamindawarCounty) BYmanda伊曼达() feud.Barony {
	return c.伊曼达Ymanda
}

func (c *扎敏达瓦尔ZamindawarCounty) BZamindawar扎敏达瓦尔() feud.Barony {
	return c.扎敏达瓦尔Zamindawar
}

var CZamindawar扎敏达瓦尔 ZamindawarCounty = &扎敏达瓦尔ZamindawarCounty{}

func init() {
	f := CZamindawar扎敏达瓦尔.(*扎敏达瓦尔ZamindawarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "854",
		Title:     "zamindawar",
		TitleName: "扎敏达瓦尔",
		TitleCode: "c_zamindawar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫吉莱Afgilai = BAfgilai阿夫吉莱
	f.阿夫吉莱Afgilai.SetParent(f)

	f.阿曼达瓦尔Amandawar = BAmandawar阿曼达瓦尔
	f.阿曼达瓦尔Amandawar.SetParent(f)

	f.比什朗格Bishlang = BBishlang比什朗格
	f.比什朗格Bishlang.SetParent(f)

	f.哈莱Khalai = BKhalai哈莱
	f.哈莱Khalai.SetParent(f)

	f.兰吉Langi = BLangi兰吉
	f.兰吉Langi.SetParent(f)

	f.伊曼达Ymanda = BYmanda伊曼达
	f.伊曼达Ymanda.SetParent(f)

	f.扎敏达瓦尔Zamindawar = BZamindawar扎敏达瓦尔
	f.扎敏达瓦尔Zamindawar.SetParent(f)

}
