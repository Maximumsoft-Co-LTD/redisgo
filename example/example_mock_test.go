package redis_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/Maximumsoft-Co-LTD/redisgo/redis"

	gomock "go.uber.org/mock/gomock"
)

// Example unit test usage with redis.MockClient (gomock style):
//
// 1. Make your code depend on redis.ClientInterface so tests can inject a mock.
// 2. Create controller and mock: ctrl := gomock.NewController(t); mock := redis.NewMockClient(ctrl)
// 3. Set expected calls with mock.EXPECT().Method(...).Return(...) or DoAndReturn(...)
// 4. Run your code; controller is finished via gomock when the test ends.

// keyStoreGet loads a key (accepts ClientInterface for testing).
func keyStoreGet(ctx context.Context, store redis.ClientInterface, key string, out interface{}) error {
	return store.Get(ctx, key, out)
}

// keyStoreSet writes a key (accepts ClientInterface for testing).
func keyStoreSet(ctx context.Context, store redis.ClientInterface, key string, ttl time.Duration, v interface{}) error {
	return store.Set(ctx, key, ttl, v)
}

func TestKeyStoreGet(t *testing.T) {
	ctx := context.Background()
	type Item struct{ Name string }

	t.Run("with mock", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := redis.NewMockClient(ctrl)

		mock.EXPECT().
			Get(ctx, "test:item", gomock.Any()).
			DoAndReturn(func(ctx context.Context, key string, obj interface{}) error {
				return json.Unmarshal([]byte(`{"name":"mocked"}`), obj)
			})

		var out Item
		err := keyStoreGet(ctx, mock, "test:item", &out)
		if err != nil {
			t.Fatalf("keyStoreGet: %v", err)
		}
		if out.Name != "mocked" {
			t.Errorf("out.Name = %q, want %q", out.Name, "mocked")
		}
	})

	t.Run("mock returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := redis.NewMockClient(ctrl)
		wantErr := errors.New("redis down")

		mock.EXPECT().
			Get(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(wantErr)

		var out struct{ X int }
		err := keyStoreGet(ctx, mock, "any", &out)
		if err != wantErr {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})
}

func TestKeyStoreSet(t *testing.T) {
	ctx := context.Background()
	key := "test:set"
	ttl := 5 * time.Second

	t.Run("with mock", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := redis.NewMockClient(ctrl)
		var setKey string
		var setTTL time.Duration
		var setPayload []byte

		mock.EXPECT().
			Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, k string, d time.Duration, obj interface{}) error {
				setKey = k
				setTTL = d
				setPayload, _ = json.Marshal(obj)
				return nil
			})

		item := struct{ Name string }{Name: "saved"}
		err := keyStoreSet(ctx, mock, key, ttl, &item)
		if err != nil {
			t.Fatalf("keyStoreSet: %v", err)
		}
		if setKey != key {
			t.Errorf("set key = %q, want %q", setKey, key)
		}
		if setTTL != ttl {
			t.Errorf("set ttl = %v, want %v", setTTL, ttl)
		}
		var decoded struct{ Name string }
		_ = json.Unmarshal(setPayload, &decoded)
		if decoded.Name != "saved" {
			t.Errorf("decoded.Name = %q, want %q", decoded.Name, "saved")
		}
	})
}

// listPushPushPop pushes two items then pops one (accepts ClientInterface for testing).
func listPushPushPop(ctx context.Context, client redis.ClientInterface, key string, v1, v2 interface{}) (popped interface{}, err error) {
	if err = client.SetList(ctx, key, v1); err != nil {
		return nil, err
	}
	if err = client.SetList(ctx, key, v2); err != nil {
		return nil, err
	}
	var out map[string]string
	if err = client.PopList(ctx, key, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func TestListPushPushPop(t *testing.T) {
	ctx := context.Background()
	key := "test:list"
	popReturn := map[string]string{"name": "first"}

	t.Run("with mock", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := redis.NewMockClient(ctrl)

		mock.EXPECT().SetList(ctx, key, gomock.Any()).Return(nil).Times(2)
		mock.EXPECT().
			PopList(ctx, key, gomock.Any()).
			DoAndReturn(func(ctx context.Context, k string, obj interface{}) error {
				data, _ := json.Marshal(popReturn)
				return json.Unmarshal(data, obj)
			})

		popped, err := listPushPushPop(ctx, mock, key, "a", "b")
		if err != nil {
			t.Fatalf("listPushPushPop: %v", err)
		}
		m, ok := popped.(map[string]string)
		if !ok {
			t.Fatalf("popped type = %T", popped)
		}
		if m["name"] != "first" {
			t.Errorf("popped[name] = %q, want %q", m["name"], "first")
		}
	})
}

func TestMockClient_expectCalls(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := redis.NewMockClient(ctrl)
	ctx := context.Background()

	mock.EXPECT().Close().Return(nil)
	mock.EXPECT().Ping(ctx).Return(nil)
	mock.EXPECT().Get(ctx, "k", nil).Return(nil)
	mock.EXPECT().Set(ctx, "k", time.Duration(0), nil).Return(nil)
	mock.EXPECT().LenList(ctx, "list").Return(int64(0), nil)

	_ = mock.Close()
	_ = mock.Ping(ctx)
	_ = mock.Get(ctx, "k", nil)
	_ = mock.Set(ctx, "k", 0, nil)
	n, _ := mock.LenList(ctx, "list")
	if n != 0 {
		t.Errorf("LenList: n=%d", n)
	}
}
