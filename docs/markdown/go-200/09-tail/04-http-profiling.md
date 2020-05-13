# http.pprof

- **Expose** les informations de **profiling** compatibles **pprof**, via une **API HTTP**.
- Il suffit d’importer le package dans son programme : `import _ "net/http/pprof"`
- `$ go tool pprof http://localhost:8080/debug/pprof/heap`
- flame graph disponible également

![center h-400](./assets/go-200/images/easy.gif)

Notes:
OFU

plein d’outils pour diagnostiquer les problèmes

comme jconsole

PRODUCTION READY
