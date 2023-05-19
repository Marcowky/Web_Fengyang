export function BarRouteGoto(router, index) {
    if (index.length == 1) {
        switch (index) {
            case "1":
                router.push("")
                break;
            case "6":
                router.push("/travel")
                break;
            default:
                console.log("wrong index1:", index)
        }
    } else if (index.length == 3) {
        switch (index[0]) {
            case "2":
                router.push({
                    path: "/info",
                    query: {
                        category: index
                    }
                })
                break;
            case "3":
                router.push({
                    path: "/attraction",
                    query: {
                        category: index
                    }
                })
                break;
            case "4":
                router.push({
                    path: "/consumption",
                    query: {
                        category: index
                    }
                })
                break;
            case "5":
                router.push({
                    path: "/blog",
                    query: {
                        category: index
                    }
                })
                break;
            default:
                console.log("wrong index3:", index)
        }

    }
}