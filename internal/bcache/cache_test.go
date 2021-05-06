package bcache_test

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/darmiel/yaxc/internal/server"
	"github.com/stretchr/testify/require"
)

// var cache = fserver.Backend.NewCache(time.Millisecond*1, time.Millisecond*2)

// var server = server.NewServer(&server.YAxCConfig{
// 	// cache: cache,
// })

// yep let's have a global cache over all tests

// note: might take some time...
func TestCacheCorruption(t *testing.T) {

	var server = server.NewServer(&server.YAxCConfig{})

	go func() {
		server.Start()
	}()

	// server.Start()

	var wg sync.WaitGroup

	INITIAL := 70

	// test.WithTestServer(t, func(s *api.Server) {
	for i := 1; i <= INITIAL; i++ {
		wg.Add(1)
		go func(t *testing.T, i int) {
			defer wg.Done()
			var wg2 sync.WaitGroup
			defer wg2.Wait()

			h := md5.New()
			key := "keyh_" + fmt.Sprint(i)
			fmt.Fprintf(h, "%s", key)
			key = fmt.Sprintf("%x", h.Sum(nil))
			val := "a"
			for it := range make([]uint64, i) {
				val = val + fmt.Sprint(it)
				key = key + "0"
			}

			_, err := url.Parse("https://" + key + ".com")
			require.NoError(t, err, "url parse error")

			// val := "val_test" + fmt.Sprintf("%06d", i)
			initial, err := server.Backend.Get(key)
			if err == nil {
				if initial != "" {
					require.Equal(t, initial, val, key+" initial already set later (not set yet)")
				}
			} else {
				// require.Nil(t, initial, key+" initial get string")
			}
			// require.False(t, ok, key+" initial get")
			// require.Nil(t, initial, key+" initial get string")

			server.Backend.Set(key, val, time.Millisecond*2)

			for ii := 1; ii <= INITIAL; ii++ {
				wg2.Add(1)

				go func(t *testing.T, ii int) {
					defer wg2.Done()

					time.Sleep(time.Microsecond * 50)

					h := md5.New()
					key := "keyh_" + fmt.Sprint(ii)
					fmt.Fprintf(h, "%s", key)
					key = fmt.Sprintf("%x", h.Sum(nil))
					// val := "val_test" + fmt.Sprintf("%06d", i)
					val := "a"
					for it := range make([]uint64, i) {
						val = val + fmt.Sprint(it)
						key = key + "0"
					}

					_, err := url.Parse("https://" + key + ".com")
					require.NoError(t, err, "url parse error")

					readVal, err := server.Backend.Get(key)

					if err == nil {
						if readVal != "" {
							require.Equal(t, val, readVal, key+" reread later")
						}

					} else {
						require.Nil(t, readVal, key+" reread later (not set yet)")
						server.Backend.Set(key, val, time.Millisecond*2)
					}

					// require.True(t, ok, key+" reread later ok")
					// require.Equal(t, expected, readVal, key+" reread later")
				}(t, ii)
			}

		}(t, i)

	}

	wg.Wait()
	// })

	// test.WithTestServer(t, func(s *api.Server) {

	// // must possible to savely get them now...
	// for i := 1; i <= INITIAL; i++ {
	// 	wg.Add(1)
	// 	go func(t *testing.T, i int) {
	// 		defer wg.Done()

	// 		h := md5.New()
	// 		key := "keyh_" + fmt.Sprint(i)
	// 		fmt.Fprintf(h, "%s", key)
	// 		key = fmt.Sprintf("%x", h.Sum(nil))
	// 		// val := "val_test" + fmt.Sprintf("%06d", i)

	// 		val := "a"
	// 		for it := range val {
	// 			val = val + string(it)
	// 			key = key + "0"
	// 		}

	// 		_, err := url.Parse("https://" + key + ".com")
	// 		require.NoError(t, err, "url parse error")

	// 		initial, err := server.Backend.Get(key)
	// 		if err != nil {
	// 			require.True(t, ok, key+" initial get")
	// 			require.Equal(t, val, initial, key+" reread 2nd loop get string")
	// 		}

	// 	}(t, i)
	// }

	// wg.Wait()

	// // })

	// LOOPNEXT := 200

	// // test.WithTestServer(t, func(s *api.Server) {

	// // must possible to savely get them now...
	// for i := 1; i <= LOOPNEXT; i++ {
	// 	wg.Add(1)
	// 	go func(t *testing.T, i int) {
	// 		defer wg.Done()

	// 		h := md5.New()
	// 		key := "keyh_" + fmt.Sprint(i)
	// 		fmt.Fprintf(h, "%s", key)
	// 		key = fmt.Sprintf("%x", h.Sum(nil))
	// 		// val := "val_test" + fmt.Sprintf("%06d", i)

	// 		val := make([]uint64, i)
	// 		for it := range val {
	// 			val[it] = uint64(65280 + it)
	// 			key = key + "0"
	// 		}

	// 		_, err := url.Parse("https://" + key + ".com")
	// 		require.NoError(t, err, "url parse error")

	// 		initial, ok := server.Backend.Get(key)

	// 		if i > INITIAL {
	// 			if ok {
	// 				require.Nil(t, initial, key+" initial get string")
	// 			}
	// 		} else {
	// 			if ok {
	// 				require.Equal(t, val, initial, key+" reread 3nd loop get string")
	// 			}
	// 		}

	// 		server.Backend.Set(key, val, time.Millisecond*2)

	// 	}(t, i)
	// }

	// wg.Wait()

	// // })

	// // test.WithTestServer(t, func(s *api.Server) {

	// // must possible to savely get them now...
	// for i := 1; i <= LOOPNEXT; i++ {
	// 	wg.Add(1)
	// 	go func(t *testing.T, i int) {
	// 		defer wg.Done()

	// 		h := md5.New()
	// 		key := "keyh_" + fmt.Sprint(i)
	// 		fmt.Fprintf(h, "%s", key)
	// 		key = fmt.Sprintf("%x", h.Sum(nil))
	// 		// val := "val_test" + fmt.Sprintf("%06d", i) + "run 2"

	// 		val := make([]uint64, i)
	// 		for it := range val {
	// 			val[it] = uint64(65280 + it)
	// 			key = key + "0"
	// 		}

	// 		_, err := url.Parse("https://" + key + ".com")
	// 		require.NoError(t, err, "url parse error")

	// 		initial, ok := server.Backend.Get(key)
	// 		if ok {
	// 			require.True(t, ok, key+" initial get")
	// 			require.Equal(t, val, initial, key+" reread 4nd loop get string")
	// 		}

	// 	}(t, i)
	// }

	// wg.Wait()

	// })
}

// func TestCacheCorruption2(t *testing.T) {

// 	var wg sync.WaitGroup
// 	LOOPNEXT := 200
// 	for i := 1; i <= LOOPNEXT; i++ {
// 		wg.Add(1)
// 		go func(t *testing.T, i int) {
// 			defer wg.Done()

// 			h := md5.New()
// 			key := "keyh_" + fmt.Sprint(i)
// 			fmt.Fprintf(h, "%s", key)
// 			key = fmt.Sprintf("%x", h.Sum(nil))
// 			// val := "val_test" + fmt.Sprintf("%06d", i) + "run 2"

// 			val := make([]uint64, i)
// 			for it := range val {
// 				val[it] = uint64(65280 + it)
// 				key = key + "0"
// 			}

// 			_, err := url.Parse("https://" + key + ".com")
// 			require.NoError(t, err, "url parse error")

// 			initial, ok := server.Backend.Get(key)
// 			require.True(t, ok, key+" initial get")
// 			require.Equal(t, val, initial, key+" reread 4nd loop get string")

// 		}(t, i)
// 	}

// 	wg.Wait()

// }
