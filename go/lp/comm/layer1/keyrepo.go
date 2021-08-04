package layer1

type KeyRepository interface {
	Set(clientID string, publicKey []byte) error
	Get(clientID string) ([]byte, error)
}

type BasicKeyRepository struct {
	keys map[string][]byte
}

func NewBasicKeyRespository() *BasicKeyRepository {
	return &BasicKeyRepository{
		keys: make(map[string][]byte),
	}
}

func (r *BasicKeyRepository) Set(clientID string, publicKey []byte) error {
	r.keys[clientID] = publicKey
	return nil
}

func (r *BasicKeyRepository) Get(clientID string) ([]byte, error) {
	publicKey := r.keys[clientID]
	return publicKey, nil
}
