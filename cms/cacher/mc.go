package cacher

type Conn interface {
	// Get retrieves a value from the cache.
	Get(key string) (val string, flags uint32, cas uint64, err error)
	// GAT (get and touch) retrieves the value associated with the key and updates
	// its expiration time.
	GAT(key string, exp uint32) (val string, flags uint32, cas uint64, err error)
	// Touch updates the expiration time on a key/value pair in the cache.
	Touch(key string, exp uint32) (cas uint64, err error)
	// Set sets a key/value pair in the cache.
	Set(key, val string, flags, exp uint32, ocas uint64) (cas uint64, err error)
	// Replace replaces an existing key/value in the cache. Fails if key doesn't
	// already exist in cache.
	Replace(key, val string, flags, exp uint32, ocas uint64) (cas uint64, err error)
	// Add adds a new key/value to the cache. Fails if the key already exists in the
	// cache.
	Add(key, val string, flags, exp uint32) (cas uint64, err error)
	// Incr increments a value in the cache. The value must be an unsigned 64bit
	// integer stored as an ASCII string. It will wrap when incremented outside the
	// range.
	Incr(key string, delta, init uint64, exp uint32, ocas uint64) (n, cas uint64, err error)
	// Decr decrements a value in the cache. The value must be an unsigned 64bit
	// integer stored as an ASCII string. It can't be decremented below 0.
	Decr(key string, delta, init uint64, exp uint32, ocas uint64) (n, cas uint64, err error)
	// Append appends the value to the existing value for the key specified. An
	// error is thrown if the key doesn't exist.
	Append(key, val string, ocas uint64) (cas uint64, err error)
	// Prepend prepends the value to the existing value for the key specified. An
	// error is thrown if the key doesn't exist.
	Prepend(key, val string, ocas uint64) (cas uint64, err error)
	// Del deletes a key/value from the cache.
	Del(key string) (err error)
	// DelCAS deletes a key/value from the cache but only if the CAS specified
	// matches the CAS in the cache.
	DelCAS(key string, cas uint64) (err error)
	// Flush flushes the cache, that is, invalidate all keys. Note, this doesn't
	// typically free memory on a memcache server (doing so compromises the O(1)
	// nature of memcache). Instead nearly all servers do lazy expiration, where
	// they don't free memory but won't return any keys to you that have expired.
	Flush(when uint32) (err error)
	// NoOp sends a No-Op message to the memcache server. This can be used as a
	// heartbeat for the server to check it's functioning fine still.
	NoOp() (err error)
	// Version gets the version of the memcached server connected to.
	Version() (ver string, err error)
	// Quit closes the connection with memcache server (nicely).
	Quit() (err error)
	// StatsWithKey returns some statistics about the memcached server. It supports
	// sending across a key to the server to select which statistics should be
	// returned.
	StatsWithKey(key string) (stats map[string]string, err error)
	// Stats returns some statistics about the memcached server.
	Stats() (stats map[string]string, err error)
	// StatsReset resets the statistics stored at the memcached server.
	StatsReset() (err error)
}
