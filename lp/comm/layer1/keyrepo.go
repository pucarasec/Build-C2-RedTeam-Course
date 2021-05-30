package layer1

type KeyRepository interface {
	Set(clientID []byte, publicKey []byte) error
	Get(clientID []byte) ([]byte, error)
}

type BasicKeyRepository struct {
	keys map[string][]byte
}

func NewBasicKeyRespository() *BasicKeyRepository {
	return &BasicKeyRepository{
		keys: make(map[string][]byte),
	}
}

func (r *BasicKeyRepository) Set(clientID []byte, publicKey []byte) error {
	r.keys[string(clientID)] = publicKey
	return nil
}

func (r *BasicKeyRepository) Get(clientID []byte) ([]byte, error) {
	publicKey := r.keys[string(clientID)]
	return publicKey, nil
}
