package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
	"time"
	"url-shortener/daos"
	"url-shortener/database/sqlite"
	"url-shortener/dtos"
)

const (
	WORKER_POOL_SZ = 12
)

var (
	dao     *daos.Dao
	service *Service
)

func TestMain(m *testing.M) {
	dao = daos.NewDao(sqlite.Open(), sqlite.OpenNonConcurrent())
	dao.CreateTables()
	service = NewService(dao)
	code := m.Run()
	dao.Destroy()
	os.Exit(code)
}

func makeShortUrl(ctx context.Context, t *testing.T) {
	startTime := time.Now()
	originalUrl := dtos.UrlToShort{
		Url:        "ya.ru",
		Expiration: time.Now().Add(time.Duration(5) * time.Minute),
		Author:     "new_user",
	}
	url, err := service.MakeShortUrl(ctx, originalUrl)
	if assert.NoError(t, err) {
		assert.Equal(t, url.Author.Name, originalUrl.Author)
		assert.Equal(t, url.Original, originalUrl.Url)
		assert.True(t, originalUrl.Expiration.Sub(url.Expiration).Abs() < time.Second)
		assert.WithinRange(t, url.Created, startTime, time.Now())
	}
}

func TestService_MakeShortUrl(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	makeShortUrl(ctx, t)
}

func TestService_GetUrlDataByShort(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()
	startTime := time.Now()
	originalUrl := dtos.UrlToShort{
		Url:        "getdatatest.ru",
		Expiration: time.Now().Add(time.Duration(5) * time.Minute),
		Author:     "new_user",
	}
	url, err := service.MakeShortUrl(ctx, originalUrl)
	assert.NoError(t, err)
	gotUrl, getErr := service.GetUrlDataByShort(url.Shorten)
	if assert.NoError(t, getErr) {
		assert.Equal(t, gotUrl.Original, originalUrl.Url)
		assert.Equal(t, gotUrl.Author.Name, originalUrl.Author)
		assert.True(t, originalUrl.Expiration.Sub(gotUrl.Expiration).Abs() < time.Second)
		assert.WithinRange(t, url.Created, startTime, time.Now())
	}
}

func TestService_GetUrlByShort(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()
	originalUrl := dtos.UrlToShort{
		Url:        "geturltest.ru",
		Expiration: time.Now().Add(time.Duration(5) * time.Minute),
		Author:     "new_user",
	}
	url, err := service.MakeShortUrl(ctx, originalUrl)
	if assert.NoError(t, err) {
		gotUrl, getErr := service.GetUrlByShort(url.Shorten)
		if assert.NoError(t, getErr) {
			assert.Equal(t, gotUrl, "geturltest.ru")
		}
	}
}

func simplestWorkerPool(b *testing.B, workers int, test func()) {
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	workerChan := make(chan struct{}, workers)
	defer close(workerChan)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			workerChan <- struct{}{}
			test()
			<-workerChan
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkService_GetUrlDataByShort(b *testing.B) {
	originalUrl := dtos.UrlToShort{
		Url:        "ya.ru",
		Expiration: time.Now().Add(time.Duration(5) * time.Minute),
		Author:     "vlad",
	}
	shortUrl, err := service.MakeShortUrl(context.TODO(), originalUrl)
	if err != nil {
		b.Error(err)
	}
	simplestWorkerPool(b, WORKER_POOL_SZ, func() {
		_, err = service.GetUrlDataByShort(shortUrl.Shorten)
		if err != nil {
			b.Error(err)
		}
	})
}

func BenchmarkService_MakeUrl(b *testing.B) {
	simplestWorkerPool(b, WORKER_POOL_SZ, func() {
		url := dtos.UrlToShort{
			Url:        "github.com/techschool/simplebank/blob/master/db/sqlc/db.go",
			Expiration: time.Now().Add(time.Duration(5) * time.Minute),
			Author:     randomFromString("vlad"),
		}
		_, err := service.MakeShortUrl(context.TODO(), url)
		if err != nil {
			b.Error(err)
		}
	})
}
