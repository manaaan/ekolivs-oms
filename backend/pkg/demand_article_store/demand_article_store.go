package demand_article_store

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/demand_store"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
	"google.golang.org/api/iterator"
)

const Collection = "demandArticles"

type Store struct {
	FirestoreClient *firestore.Client
}

type StoreDemandArticle struct {
	demand_api.Article
}

func (s Store) GetArticles(ctx context.Context, demand *demand_api.Demand) ([]*demand_api.Article, error) {
	var articles []*demand_api.Article

	iter := s.FirestoreClient.Collection(demand_store.Collection).Doc(demand.ID).Collection(Collection).Documents(ctx)
	defer iter.Stop()
	for {
		dsnap, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		var prod demand_api.Article
		if err := dsnap.DataTo(&prod); err != nil {
			return nil, err
		}
		articles = append(articles, &prod)
	}

	return articles, nil
}

func (s Store) CreateOrUpdateDemandArticle(ctx context.Context, demand *demand_api.Demand, article *demand_api.Article) (*demand_api.Article, error) {
	dr := prepToCreateOrUpdateDemandArticle(s.FirestoreClient, demand.ID, article)
	if _, err := dr.Set(ctx, article); err != nil {
		return nil, err
	}

	return article, nil
}

func (s Store) CreateOrUpdateDemandArticleWithTx(tx *firestore.Transaction, demandId string, article *demand_api.Article) (*demand_api.Article, error) {
	dr := prepToCreateOrUpdateDemandArticle(s.FirestoreClient, demandId, article)
	if err := tx.Set(dr, article); err != nil {
		return nil, err
	}

	return article, nil
}

func prepToCreateOrUpdateDemandArticle(firestoreClient *firestore.Client, demandId string, article *demand_api.Article) *firestore.DocumentRef {
	if len(article.ID) == 0 {
		article.ID = firestoreClient.Collection(Collection).NewDoc().ID
	}

	dr := firestoreClient.Collection(demand_store.Collection).Doc(demandId).Collection(Collection).Doc(article.ID)
	return dr
}

func (s Store) DeleteDemandArticle(ctx context.Context, demand *demand_api.Demand, id string) error {
	if _, err := s.FirestoreClient.Collection(demand_store.Collection).Doc(demand.ID).Collection(Collection).Doc(id).Delete(ctx); err != nil {
		return err
	}

	return nil
}
