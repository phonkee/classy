# Classy

Class based views, inspired by django class based views. The functionality is simple yet powerful.

## API

    classy.Debug()

    classy.Use(middleware1, middleware2).Name("api:{name}").Register(
        router,
        classy.New(&ProductDetailView{}).Use(middleware3),
        classy.New(&ProductApproveView{}).Path("/approve").Name("approve"),
    )

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

    // set response as not allowed
    notAllowed := response.New()

    classy.Path("/api").MethodNotAllowed(notAllowed).Register(
        router,
        classy.New(&ProductDetailView{}).Path("/product/"),
        classy.New(&ProductApproveView{}).Path("/product/approve").Debug(),
    )

    classy.Path("/api").Register(
        router,
        classy.New(&ProductDetailView{}).Path("/product/").MethodNotAllowed(notAllowed),
        classy.New(&ProductApproveView{}).Path("/product/approve").Debug().MethodNotAllowed(notAllowed),
    )

Every view needs to have Routes method that returns mapping:

    func (l View) Routes() map[string]Mapping {
        return map[string]Mapping {
            "/": NewMapping(
                []string("GET": "List"},
                []string("POST": "Post"},
            )
        }
    }

If you embed multiple views you can use shorthand to merge routes:

    return MultiRoutes().
        Add(Detail.Routes(), "{name}_detail").
        Add(List.Routes(), "{name}_list").
        Get()
    