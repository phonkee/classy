# Classy

Class based views, inspired by django class based views. The functionality is simple yet powerful.

## API

```go
    // enable debug for whole classy
    classy.Debug()

    classy.Use(middleware1, middleware2).Name("api:{name}").Register(
        router,
        classy.New(&ProductDetailView{}).Use(middleware3),
        classy.New(&ProductApproveView{}).Path("/approve").Name("approve"),
    )

    // lets register some views
    classy.Register(
        router,
        classy.New(&ProductDetailView{}).Path("/product/"),
        classy.New(&ProductApproveView{}).Path("/product/approve").Debug(),
    )

    classy.Path("/api").Register(
        router,
        classy.New(&ProductDetailView{}).Path("/product/"),
        classy.New(&ProductApproveView{}).Path("/product/approve").Debug(),
    )

    // set response as not allowed (TODO)
    na := response.New(http.StatusMethodNotAllowed)

    classy.Path("/api").MethodNotAllowed(na).Register(
        router,
        classy.New(&ProductDetailView{}).Path("/product/"),
        classy.New(&ProductApproveView{}).Path("/product/approve").Debug(),
    )

    // method not allowed (TODO)
    classy.Path("/api").Register(
        router,
        classy.New(&ProductDetailView{}).Path("/product/").MethodNotAllowed(na),
        classy.New(&ProductApproveView{}).Path("/product/approve").Debug().MethodNotAllowed(na),
    )
```

Every view needs to have Routes method that returns mapping:

```go
    func (l View) Routes() map[string]Mapping {
        return map[string]Mapping {
            "/": NewMapping(
                []string("GET": "List"},
                []string("POST": "Post"},
            )
        }
    }
```

If you embed multiple views you can use shorthand to merge routes:

```go
    return MultiRoutes().
        Add(Detail.Routes(), "{name}_detail").
        Add(List.Routes(), "{name}_list").
        Get()
```
    