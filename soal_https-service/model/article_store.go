package model

type ArticleStoreInMemory struct {
	ArticleMap []Article
}

func NewArticleStoreInMemory() *ArticleStoreInMemory {
	return &ArticleStoreInMemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "Membuat Website", Body: "ini body article"},
		},
	}
}

func (store *ArticleStoreInMemory) Save(article *Article) error {
	lastID := len(store.ArticleMap)

	article.ID = lastID + 1

	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}
