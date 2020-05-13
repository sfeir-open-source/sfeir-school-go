<!-- .slide: class="with-code" -->

# Accès aux données

## Les tests DAO

```go
package dao_test

func TestDAOMongo(t *testing.T) {
	// get config
	config := os.Getenv("DB_HOST")

    // build DAO
	daoMongo, err := dao.GetTaskDAO(config, dao.DAOMongo)
	if err != nil {
		t.Error(err)
	}

    //...
}
```

<!-- .element style="font-size:0.6em; line-height: 1em" -->

```go
package dao_test

func TestDAOMock(t *testing.T) {
	daoMock, err := dao.GetTaskDAO("", dao.DAOMock)
	if err != nil {
		t.Error(err)
	}

    //...
}

```

<!-- .element style="font-size:0.6em; line-height: 1em" -->

Notes:
OFU

TODO: implémenter le test DAO Mock
