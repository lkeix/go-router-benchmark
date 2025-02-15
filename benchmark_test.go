package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServeHTTP(b *testing.B, r route, router http.Handler) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, r.reqPath, nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(rec, req)
		if rec.Code != 200 {
			panic(fmt.Sprintf("Request failed. path: %v request path:%v", r.path, r.reqPath))
		}
	}
}

func benchmark(b *testing.B, r route, router http.Handler) {
	testServeHTTP(b, r, router)
}

// net/http#ServeMux
// https://pkg.go.dev/net/http#ServeMux
func BenchmarkStaticRoutesRootServeMux(b *testing.B) {
	router := loadServeMux(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1ServeMux(b *testing.B) {
	router := loadServeMux(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5ServeMux(b *testing.B) {
	router := loadServeMux(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10ServeMux(b *testing.B) {
	router := loadServeMux(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

// bmf-san/goblin
// https://github.com/bmf-san/goblin
func BenchmarkStaticRoutesRootGoblin(b *testing.B) {
	router := loadGoblin(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Goblin(b *testing.B) {
	router := loadGoblin(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Goblin(b *testing.B) {
	router := loadGoblin(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Goblin(b *testing.B) {
	router := loadGoblin(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1Goblin(b *testing.B) {
	router := loadGoblin(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5Goblin(b *testing.B) {
	router := loadGoblin(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10Goblin(b *testing.B) {
	router := loadGoblin(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// julienschmidt/httprouter
// https://github.com/julienschmidt/httprouter
func BenchmarkStaticRoutesRootHTTPRouter(b *testing.B) {
	router := loadHTTPRouter(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// go-chi/chi
// https://github.com/go-chi/chi
func BenchmarkStaticRoutesRootChi(b *testing.B) {
	router := loadChi(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Chi(b *testing.B) {
	router := loadChi(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Chi(b *testing.B) {
	router := loadChi(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Chi(b *testing.B) {
	router := loadChi(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamBracketRoutes1Chi(b *testing.B) {
	router := loadChi(pathParamBracketRoutes1)
	benchmark(b, pathParamBracketRoutes1, router)
}

func BenchmarkPathParamBracketRoutes5Chi(b *testing.B) {
	router := loadChi(pathParamBracketRoutes5)
	benchmark(b, pathParamBracketRoutes5, router)
}

func BenchmarkPathParamBracketRoutes10Chi(b *testing.B) {
	router := loadChi(pathParamBracketRoutes10)
	benchmark(b, pathParamBracketRoutes10, router)
}

// gin-gonic/gin
// https://github.com/gin-gonic/gin
func BenchmarkStaticRoutesRootGin(b *testing.B) {
	router := loadGin(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Gin(b *testing.B) {
	router := loadGin(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Gin(b *testing.B) {
	router := loadGin(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Gin(b *testing.B) {
	router := loadGin(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1Gin(b *testing.B) {
	router := loadGin(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5Gin(b *testing.B) {
	router := loadGin(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10Gin(b *testing.B) {
	router := loadGin(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// uptrace/bunrouter
// https://github.com/uptrace/bunrouter
func BenchmarkStaticRoutesRootBunRouter(b *testing.B) {
	router := loadBunRouter(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1BunRouter(b *testing.B) {
	router := loadBunRouter(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5BunRouter(b *testing.B) {
	router := loadBunRouter(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10BunRouter(b *testing.B) {
	router := loadBunRouter(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1BunRouter(b *testing.B) {
	router := loadBunRouter(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5BunRouter(b *testing.B) {
	router := loadBunRouter(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10BunRouter(b *testing.B) {
	router := loadBunRouter(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// dimfeld/httptreemux
// https://github.com/dimfeld/httptreemux
func BenchmarkStaticRoutesRootHTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// beego/mux
// https://github.com/beego/mux
func BenchmarkStaticRoutesRootBeegoMux(b *testing.B) {
	router := loadBeegoMux(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1BeegoMux(b *testing.B) {
	router := loadBeegoMux(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5BeegoMux(b *testing.B) {
	router := loadBeegoMux(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10BeegoMux(b *testing.B) {
	router := loadBeegoMux(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1BeegoMux(b *testing.B) {
	router := loadBeegoMux(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5BeegoMux(b *testing.B) {
	router := loadBeegoMux(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10BeegoMux(b *testing.B) {
	router := loadBeegoMux(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// gorilla/mux
// https://github.com/gorilla/mux
func BenchmarkStaticRoutesRootGorillaMux(b *testing.B) {
	router := loadGorillaMux(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1GorillaMux(b *testing.B) {
	router := loadGorillaMux(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5GorillaMux(b *testing.B) {
	router := loadGorillaMux(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10GorillaMux(b *testing.B) {
	router := loadGorillaMux(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamBracketRoutes1GorillaMux(b *testing.B) {
	router := loadGorillaMux(pathParamBracketRoutes1)
	benchmark(b, pathParamBracketRoutes1, router)
}

func BenchmarkPathParamBracketRoutes5GorillaMux(b *testing.B) {
	router := loadGorillaMux(pathParamBracketRoutes5)
	benchmark(b, pathParamBracketRoutes5, router)
}

func BenchmarkPathParamBracketRoutes10GorillaMux(b *testing.B) {
	router := loadGorillaMux(pathParamBracketRoutes10)
	benchmark(b, pathParamBracketRoutes10, router)
}

// nissy/bon
// https://github.com/nissy/bon
func BenchmarkStaticRoutesRootBon(b *testing.B) {
	router := loadBon(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Bon(b *testing.B) {
	router := loadBon(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Bon(b *testing.B) {
	router := loadBon(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Bon(b *testing.B) {
	router := loadBon(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1Bon(b *testing.B) {
	router := loadBon(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5Bon(b *testing.B) {
	router := loadBon(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10Bon(b *testing.B) {
	router := loadBon(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// naoina/denco
// https://github.com/naoina/denco
func BenchmarkStaticRoutesRootDenco(b *testing.B) {
	router := loadDenco(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Denco(b *testing.B) {
	router := loadDenco(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Denco(b *testing.B) {
	router := loadDenco(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Denco(b *testing.B) {
	router := loadDenco(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1Denco(b *testing.B) {
	router := loadDenco(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5Denco(b *testing.B) {
	router := loadDenco(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10Denco(b *testing.B) {
	router := loadDenco(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// labstack/echo
// https://github.com/labstack/echo
func BenchmarkStaticRoutesRootEcho(b *testing.B) {
	router := loadEcho(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Echo(b *testing.B) {
	router := loadEcho(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Echo(b *testing.B) {
	router := loadEcho(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Echo(b *testing.B) {
	router := loadEcho(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1Echo(b *testing.B) {
	router := loadEcho(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5Echo(b *testing.B) {
	router := loadEcho(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10Echo(b *testing.B) {
	router := loadEcho(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// gocraft/web
// https://github.com/gocraft/web
func BenchmarkStaticRoutesRootGocraftWeb(b *testing.B) {
	router := loadGocraftWeb(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}

// vardius/gorouter
// https://github.com/vardius/gorouter
func BenchmarkStaticRoutesRootGorouter(b *testing.B) {
	router := loadGorouter(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Gorouter(b *testing.B) {
	router := loadGorouter(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Gorouter(b *testing.B) {
	router := loadGorouter(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Gorouter(b *testing.B) {
	router := loadGorouter(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamBracketRoutes1Gorouter(b *testing.B) {
	router := loadGorouter(pathParamBracketRoutes1)
	benchmark(b, pathParamBracketRoutes1, router)
}

func BenchmarkPathParamBracketRoutes5Gorouter(b *testing.B) {
	router := loadGorouter(pathParamBracketRoutes5)
	benchmark(b, pathParamBracketRoutes5, router)
}

func BenchmarkPathParamBracketRoutes10Gorouter(b *testing.B) {
	router := loadGorouter(pathParamBracketRoutes10)
	benchmark(b, pathParamBracketRoutes10, router)
}

// go-ozzo/ozzo-routing
// https://github.com/go-ozzo/ozzo-routing
func BenchmarkStaticRoutesRootOzzoRouting(b *testing.B) {
	router := loadOzzoRouting(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamInequalitySignRoutes1OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(pathParamInequalitySignRoutes1)
	benchmark(b, pathParamInequalitySignRoutes1, router)
}

func BenchmarkPathParamInequalitySignRoutes5OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(pathParamInequalitySignRoutes5)
	benchmark(b, pathParamInequalitySignRoutes5, router)
}

func BenchmarkPathParamInequalitySignRoutes10OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(pathParamInequalitySignRoutes10)
	benchmark(b, pathParamInequalitySignRoutes10, router)
}

// n9te9 router
// https://github.com/lkeix/techbook13-sample
func BenchmarkStaticRoutesRootN9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1ON9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5ON9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10N9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamColonRoutes1N9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(pathParamColonRoutes1)
	benchmark(b, pathParamColonRoutes1, router)
}

func BenchmarkPathParamColonRoutes5N9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(pathParamColonRoutes5)
	benchmark(b, pathParamColonRoutes5, router)
}

func BenchmarkPathParamColonRoutes10N9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(pathParamColonRoutes10)
	benchmark(b, pathParamColonRoutes10, router)
}
