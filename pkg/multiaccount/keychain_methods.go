package multiaccount

func (k *profileKeychain) Get(key string) ([]byte, error) {
	return k.base.Get(k.transform(key))
}

func (k *profileKeychain) Set(key string, data []byte) error {
	return k.base.Set(k.transform(key), data)
}

func (k *profileKeychain) Remove(key string) error {
	return k.base.Remove(k.transform(key))
}
